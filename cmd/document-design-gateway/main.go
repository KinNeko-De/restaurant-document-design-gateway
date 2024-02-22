package main

import (
	"os"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation/health"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation/logger"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/server"
)

func main() {
	logger.SetLogLevel(logger.LogLevel)
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

	httpServerStarted := make(chan struct{})
	httpServerStopped := make(chan struct{})
	grpcServerStarted := make(chan struct{})
	grpcServerStopped := make(chan struct{})
	go server.StartHttpServer(httpServerStarted, httpServerStopped, ":8080")
	go server.StartGrpcServer(grpcServerStarted, grpcServerStopped)

	<-grpcServerStarted
	<-httpServerStarted
	health.Ready()

	<-grpcServerStopped
	<-httpServerStopped
	logger.Logger.Info().Msg("Application stopped.")
	os.Exit(0)
}
