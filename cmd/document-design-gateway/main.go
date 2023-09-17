package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation"
)

func main() {
	operation.SetDefaultLoggingLevel()

	documentServiceConfigError := document.ReadConfig()
	if documentServiceConfigError != nil {
		operation.Logger.Fatal().Err(documentServiceConfigError).Msg("Failed to read document service config")
	}

	oauthConfigError := oauth.ReadConfig()
	if oauthConfigError != nil {
		operation.Logger.Fatal().Err(oauthConfigError).Msg("Failed to read github oauth config")
	}

	StartHttpServer()
}

func StartHttpServer() {
	router := setupRouter()
	configRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		operation.Logger.Fatal().Err(err).Msg("Failed to start http server")
	}
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(operation.GinLogger())
	router.Use(gin.Recovery())
	return router
}

func configRoutes(router *gin.Engine) {
	authorized := router.Group("/")
	authorized.Use(oauth.GithubOAuth())
	authorized.GET("/document/preview/demo", document.GeneratePreviewDemo)
	authorized.POST("/document/preview", document.GeneratePreview)
}
