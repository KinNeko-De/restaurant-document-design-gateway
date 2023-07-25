package main

import (
	"log"
	"os"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kinneko-de/restaurant-document-design-gateway/build"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
)

func main() {
	setPreviewGatewayConfig()

	log.Println("Version " + build.Version)
	router := setupRouter()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/document/preview", document.GeneratePreview)
	return router
}

func setPreviewGatewayConfig() {
	connection, err := loadApiDocumentServiceConfig()
	if err != nil {
		log.Fatal(err)
	}
	document.ApiDocumentService = connection
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
