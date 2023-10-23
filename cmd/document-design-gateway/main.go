package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation"
)

func main() {
	operation.SetDefaultLoggingLevel()
	operation.Logger.Info().Msg("Starting application.")

	documentServiceConfigError := document.ReadConfig()
	if documentServiceConfigError != nil {
		operation.Logger.Fatal().Err(documentServiceConfigError).Msg("Failed to read document service config")
	}

	oauthConfigError := oauth.ReadConfig()
	if oauthConfigError != nil {
		operation.Logger.Fatal().Err(oauthConfigError).Msg("Failed to read github oauth config")
	}

	httpServerStop := make(chan struct{})
	go StartHttpServer(httpServerStop, ":8080")

	<-httpServerStop
	operation.Logger.Info().Msg("Application stopped.")
	os.Exit(0)
}

func StartHttpServer(httpServerStop chan struct{}, port string) {
	router := setupRouter()
	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)
	operation.Logger.Debug().Msgf("starting http server")

	server := &http.Server{Addr: port, Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			operation.Logger.Fatal().Err(err).Msg("Failed to start http server")
		}
	}()

	stop := <-gracefulStop
	if err := server.Shutdown(context.Background()); err != nil {
		operation.Logger.Error().Err(err).Msg("Failed to shutdown http server")
	}
	operation.Logger.Debug().Msgf("http server stopped. Received signal %s", stop)
	close(httpServerStop)
}

func setupRouter() *gin.Engine {
	router := createRouter()
	configRoutes(router)
	return router
}

func createRouter() *gin.Engine {
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
