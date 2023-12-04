package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation/logger"
)

func SetupRouter() *gin.Engine {
	router := createRouter()
	configRoutes(router)
	return router
}

func createRouter() *gin.Engine {
	router := gin.New()
	router.Use(logger.GinLogger())
	router.Use(gin.Recovery())
	return router
}

func configRoutes(router *gin.Engine) {
	authorized := router.Group("/")
	authorized.Use(oauth.GithubOAuth())
	authorized.GET("/document/preview/demo", document.GeneratePreviewDemo)
	authorized.POST("/document/preview", document.GeneratePreview)
}
