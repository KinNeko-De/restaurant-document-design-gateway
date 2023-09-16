package document

import (
	"errors"
	"os"

	apiRestaurantDocument "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const hostEnv = "DOCUMENTGENERATESERVICE_HOST"
const portEnv = "DOCUMENTGENERATESERVICE_PORT"

var (
	apiDocumentServiceUrl string = "set_by_read_config"
)

type DocumentServiceGateway interface {
	CreateDocumentServiceClient() (apiRestaurantDocument.DocumentServiceClient, error)
}

type DocumentServiceGateKeeper struct {
}

func (DocumentServiceGateKeeper) CreateDocumentServiceClient() (apiRestaurantDocument.DocumentServiceClient, error) {
	connection, dialError := grpc.Dial(apiDocumentServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if dialError != nil {
		return nil, dialError
	}
	client := apiRestaurantDocument.NewDocumentServiceClient(connection)
	return client, nil
}

func ReadConfig() error {
	connection, err := loadApiDocumentServiceConfig()
	if err != nil {
		return err
	}
	apiDocumentServiceUrl = connection

	return nil
}

func loadApiDocumentServiceConfig() (string, error) {
	host, found := os.LookupEnv(hostEnv)
	if !found {
		return "", errors.New("service host to generate documents is not configured. Expect environment variable " + hostEnv)
	}
	port, found := os.LookupEnv(portEnv)
	if !found {
		return "", errors.New("service port to generate documents is not configured. Expect environment variable " + portEnv)
	}

	return host + ":" + port, nil
}
