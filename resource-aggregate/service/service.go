package service

import (
	"context"
	"crypto/tls"
	"os"
	"os/signal"
	"sync"
	"syscall"

	clientAS "github.com/plgd-dev/cloud/authorization/client"
	pbAS "github.com/plgd-dev/cloud/authorization/pb"
	"github.com/plgd-dev/cloud/resource-aggregate/pb"
	cqrsEventBus "github.com/plgd-dev/cqrs/eventbus"
	cqrsEventStore "github.com/plgd-dev/cqrs/eventstore"
	cqrsMaintenance "github.com/plgd-dev/cqrs/eventstore/maintenance"
	"github.com/plgd-dev/kit/log"
	kitNetGrpc "github.com/plgd-dev/kit/net/grpc"
	"github.com/plgd-dev/kit/security/jwt"
	"github.com/plgd-dev/kit/security/oauth/manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type EventStore interface {
	cqrsEventStore.EventStore
	cqrsMaintenance.EventStore
	GetInstanceId(ctx context.Context, resourceId string) (int64, error)
	RemoveInstanceId(ctx context.Context, instanceId int64) error
}

//Server handle HTTP request
type Server struct {
	server             *kitNetGrpc.Server
	cfg                Config
	handler            *RequestHandler
	sigs               chan os.Signal
	authConn           *grpc.ClientConn
	userDevicesManager *clientAS.UserDevicesManager
}

type ClientCertManager = interface {
	GetClientTLSConfig() *tls.Config
}

type ServerCertManager = interface {
	GetServerTLSConfig() *tls.Config
}

// New creates new Server with provided store and publisher.
func New(config Config, clientCertManager ClientCertManager, serverCertManager ServerCertManager, eventStore EventStore, publisher cqrsEventBus.Publisher) *Server {
	dialTLSConfig := clientCertManager.GetClientTLSConfig()
	listenTLSConfig := serverCertManager.GetServerTLSConfig()

	auth := NewAuth(config.JwksURL, dialTLSConfig)
	grpcServer, err := kitNetGrpc.NewServer(config.Config.Addr, grpc.Creds(credentials.NewTLS(listenTLSConfig)), auth.Stream(), auth.Unary())
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	oauthMgr, err := manager.NewManagerFromConfiguration(config.OAuth, dialTLSConfig)
	if err != nil {
		log.Fatalf("cannot create oauth manager: %w", err)
	}

	asConn, err := grpc.Dial(config.AuthServerAddr, grpc.WithTransportCredentials(credentials.NewTLS(dialTLSConfig)), grpc.WithPerRPCCredentials(kitNetGrpc.NewOAuthAccess(oauthMgr.GetToken)))
	if err != nil {
		log.Fatalf("cannot connect to authorization server: %w", err)
	}
	authClient := pbAS.NewAuthorizationServiceClient(asConn)

	userDevicesManager := clientAS.NewUserDevicesManager(userDevicesChanged, authClient, config.UserDevicesManagerTickFrequency, config.UserDevicesManagerExpiration, func(err error) { log.Errorf("resource-aggregate: error occurs during receiving devices: %v", err) })
	requestHandler := NewRequestHandler(config, eventStore, publisher, func(ctx context.Context, userID, deviceID string) (bool, error) {
		devices, err := userDevicesManager.GetUserDevices(ctx, userID)
		if err != nil {
			return false, err
		}
		for _, id := range devices {
			if id == deviceID {
				return true, nil
			}
		}
		devices, err = userDevicesManager.UpdateUserDevices(ctx, userID)
		if err != nil {
			return false, err
		}
		for _, id := range devices {
			if id == deviceID {
				return true, nil
			}
		}
		return false, nil
	})
	pb.RegisterResourceAggregateServer(grpcServer.Server, requestHandler)

	server := Server{
		server:             grpcServer,
		cfg:                config,
		handler:            requestHandler,
		sigs:               make(chan os.Signal, 1),
		authConn:           asConn,
		userDevicesManager: userDevicesManager,
	}

	return &server
}

func (s *Server) serveWithHandlingSignal(serve func() error) error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)
	go func(s *Server) {
		defer wg.Done()
		err = serve()
	}(s)

	signal.Notify(s.sigs,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-s.sigs

	s.server.Stop()
	wg.Wait()
	s.userDevicesManager.Close()
	s.authConn.Close()

	return err
}

// Serve serve starts the service's HTTP server and blocks.
func (s *Server) Serve() error {
	return s.serveWithHandlingSignal(s.server.Serve)
}

// Shutdown ends serving
func (s *Server) Shutdown() {
	s.sigs <- syscall.SIGTERM
}

func NewAuth(jwksUrl string, tls *tls.Config) kitNetGrpc.AuthInterceptors {
	interceptor := kitNetGrpc.ValidateJWT(jwksUrl, tls, func(ctx context.Context, method string) kitNetGrpc.Claims {
		return jwt.NewScopeClaims()
	})
	return kitNetGrpc.MakeAuthInterceptors(func(ctx context.Context, method string) (context.Context, error) {
		ctx, err := interceptor(ctx, method)
		if err != nil {
			log.Errorf("auth interceptor: %v", err)
			return ctx, err
		}
		userID, err := kitNetGrpc.UserIDFromMD(ctx)
		if err != nil {
			userID, err = kitNetGrpc.UserIDFromTokenMD(ctx)
			if err == nil {
				ctx = kitNetGrpc.CtxWithIncomingUserID(ctx, userID)
			}
		}
		if err != nil {
			log.Errorf("auth cannot get userID: %v", err)
			return ctx, err
		}
		return kitNetGrpc.CtxWithUserID(ctx, userID), nil
	})
}
