package service_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"
	"testing"

	"github.com/plgd-dev/cloud/http-gateway/uri"
	testCfg "github.com/plgd-dev/cloud/test/config"

	"github.com/gorilla/websocket"
	authTest "github.com/plgd-dev/cloud/authorization/provider"
	"github.com/plgd-dev/cloud/grpc-gateway/client"
	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	"github.com/plgd-dev/cloud/http-gateway/service"
	"github.com/plgd-dev/cloud/http-gateway/test"
	cloudTest "github.com/plgd-dev/cloud/test"
	"github.com/plgd-dev/kit/codec/json"
	kitNetGrpc "github.com/plgd-dev/kit/net/grpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TestDeviceResourcesObservation(t *testing.T) {
	deviceID := cloudTest.MustFindDeviceByName(cloudTest.TestDeviceName)
	//set up
	ctx, cancel := context.WithTimeout(context.Background(), 2*test.TestTimeout)
	defer cancel()
	ctx = kitNetGrpc.CtxWithToken(ctx, authTest.UserToken)
	tearDown := cloudTest.SetUp(ctx, t)
	defer tearDown()
	webTearDown := test.SetUp(t)
	defer webTearDown()

	//onboard
	conn, err := grpc.Dial(testCfg.GRPC_HOST, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		RootCAs: cloudTest.GetRootCertificatePool(t),
	})))
	require.NoError(t, err)
	c := pb.NewGrpcGatewayClient(conn)
	defer conn.Close()
	deviceID, shutdownDevSim := cloudTest.OnboardDevSim(ctx, t, c, deviceID, testCfg.GW_HOST, cloudTest.GetAllBackendResourceLinks())

	// create web socket connection
	wsConn := webSocketConnection(t, GetDeviceResourcesObservationUri(deviceID))
	defer closeWebSocketConnection(t, wsConn)

	//read messages
	received := sync.Map{}
	go readMessage(t, wsConn, &received)

	//ofboard
	shutdownDevSim()
}

func readMessage(t *testing.T, conn *websocket.Conn, messages *sync.Map) {
	for {
		tp, message, err := conn.ReadMessage()
		if tp == websocket.CloseMessage {
			return
		}
		if err != nil {
			return
		}
		var event service.DeviceResourceObservationEvent
		err = json.Decode(message, &event)
		require.NoError(t, err)
		if event.Event == service.ToDeviceResourcesObservationEvent(client.DeviceResourcesObservationEvent_ADDED) {
			messages.Store("added", event)
		} else if event.Event == service.ToDeviceResourcesObservationEvent(client.DeviceResourcesObservationEvent_REMOVED) {
			messages.Delete("deleted")
		}
	}
}

func GetDeviceResourcesObservationUri(deviceID string) string {
	return fmt.Sprintf("wss://localhost:%d%s/%s", test.HTTP_GW_Port, uri.WSDevices, deviceID)
}
