package document

import (
	"errors"
	"log"
	"os"

	apiRestaurantDocument "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	apiDocumentServiceUrl string = "set_by_init"
)

func init() {
	readConfig()
}

type DocumentServiceGateway interface {
	CreateDocumentServiceClient() (apiRestaurantDocument.DocumentServiceClient, error)
}

type DocumentServiceGateKeeper struct {
}

func (DocumentServiceGateKeeper) CreateDocumentServiceClient() (apiRestaurantDocument.DocumentServiceClient, error) {
	connection, dialError := grpc.Dial(apiDocumentServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if dialError != nil {
		return nil, dialError;
	}
	client := apiRestaurantDocument.NewDocumentServiceClient(connection)
	return client, nil
}

func readConfig() {
	connection, err := loadApiDocumentServiceConfig()
	if err != nil {
		log.Fatal(err)
	}
	apiDocumentServiceUrl = connection
}

func loadApiDocumentServiceConfig() (string, error){
	host, found := os.LookupEnv("DOCUMENTGENERATESERVICE_HOST")
	if(!found) {
		return "", errors.New("service host to generate documents is not configured. Expect environment variable DOCUMENTSERVICE_HOST")
	}
	port, found := os.LookupEnv("DOCUMENTGENERATESERVICE_PORT")
	if(!found) {
		return "", errors.New("service port to generate documents is not configured. Expect environment variable DOCUMENTSERVICE_PORT")
	}

	return host + ":" + port, nil
}