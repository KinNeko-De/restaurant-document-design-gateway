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
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation/logger"
)

func main() {
	logger.SetInfoLogLevel()
	logger.Logger.Info().Msg("Starting application.")

	documentServiceConfigError := document.ReadConfig()
	if documentServiceConfigError != nil {
		logger.Logger.Error().Err(documentServiceConfigError).Msg("Failed to read document service config")
		os.Exit(40)
	}

	oauthConfigError := oauth.ReadConfig()
	if oauthConfigError != nil {
		logger.Logger.Error().Err(oauthConfigError).Msg("Failed to read github oauth config")
		os.Exit(41)
	}

	httpServerStop := make(chan struct{})
	go startHttpServer(httpServerStop, ":8080")

	<-httpServerStop
	logger.Logger.Info().Msg("Application stopped.")
	os.Exit(0)
}

func startHttpServer(httpServerStop chan struct{}, port string) {
	router := setupRouter()
	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)
	logger.Logger.Debug().Msgf("starting http server")

	server := &http.Server{Addr: port, Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Error().Err(err).Msg("Failed to start http server")
			os.Exit(50)
		}
	}()

	stop := <-gracefulStop
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Logger.Error().Err(err).Msg("Failed to shutdown http server")
	}
	logger.Logger.Debug().Msgf("http server stopped. Received signal %s", stop)
	close(httpServerStop)
}

func setupRouter() *gin.Engine {
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
