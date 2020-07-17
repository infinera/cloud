package refImpl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-ocf/cloud/cloud2cloud-connector/service"
	storeMongodb "github.com/go-ocf/cloud/cloud2cloud-connector/store/mongodb"
	"github.com/go-ocf/kit/log"
	"github.com/go-ocf/kit/security/certManager"
)

type Config struct {
	Log              log.Config `envconfig:"LOG"`
	Service          service.Config
	Dial             certManager.Config `envconfig:"DIAL"`
	Listen           certManager.Config `envconfig:"LISTEN"`
	ListenWithoutTLS bool               `envconfig:"LISTEN_WITHOUT_TLS"`
	StoreMongoDB     storeMongodb.Config
}

//String return string representation of Config
func (c Config) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("config: \n%v\n", string(b))
}

func Init(config Config) (*service.Server, error) {
	log.Setup(config.Log)
	log.Info(config.String())
	dialCertManager, err := certManager.NewCertManager(config.Dial)
	if err != nil {
		return nil, fmt.Errorf("cannot create dial cert manager %v", err)
	}
	dialTLSConfig := dialCertManager.GetClientTLSConfig()

	store, err := storeMongodb.NewStore(context.Background(), config.StoreMongoDB, storeMongodb.WithTLS(dialTLSConfig))
	if err != nil {
		return nil, fmt.Errorf("cannot create mongodb store %v", err)
	}

	var listenCertManager certManager.CertManager
	if !config.ListenWithoutTLS {
		listenCertManager, err = certManager.NewCertManager(config.Listen)
		if err != nil {
			return nil, fmt.Errorf("cannot create listen cert manager %v", err)
		}
	}

	return service.New(config.Service, dialCertManager, listenCertManager, store), nil
}