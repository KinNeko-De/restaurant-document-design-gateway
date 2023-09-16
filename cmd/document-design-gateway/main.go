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
	err := document.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	authorized := router.Group("/")
	authorized.Use(oauth.GithubOAuth())
	authorized.GET("/document/preview/demo", document.GeneratePreviewDemo)
	authorized.POST("/document/preview", document.GeneratePreview)

	return router
}
