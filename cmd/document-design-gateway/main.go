package main

import (
	"os"

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
	go server.StartHttpServer(httpServerStop, ":8080")

	<-httpServerStop
	logger.Logger.Info().Msg("Application stopped.")
	os.Exit(0)
}
