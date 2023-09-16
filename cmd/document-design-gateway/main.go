package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kinneko-de/restaurant-document-design-gateway/build"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
)

func main() {
	log.Println("Version " + build.Version)
	router := setupRouter()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	documentServiceConfigError := document.ReadConfig()
	if documentServiceConfigError != nil {
		log.Fatal(documentServiceConfigError)
	}
	oauthConfigError := oauth.ReadConfig()
	if oauthConfigError != nil {
		log.Fatal(oauthConfigError)
	}

	authorized := router.Group("/")
	authorized.Use(oauth.GithubOAuth())
	authorized.GET("/document/preview/demo", document.GeneratePreviewDemo)
	authorized.POST("/document/preview", document.GeneratePreview)

	return router
}
