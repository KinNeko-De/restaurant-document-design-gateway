package server

import (
	"context"
	"net/http"
	"os"

	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation/logger"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/router"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/server/shutdown"
)

func StartHttpServer(httpServerStarted chan struct{}, httpServerStopped chan struct{}) {
	port := ":8080" // todo load from env, move os.exit up to here and refactor tests

	startHttpServer(httpServerStarted, httpServerStopped, port)
}

func startHttpServer(httpServerStarted chan struct{}, httpServerStopped chan struct{}, port string) {
	router := router.SetupRouter()
	var gracefulStop = shutdown.CreateGracefulStop()
	logger.Logger.Debug().Msg("starting http server")

	server := &http.Server{Addr: port, Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Error().Err(err).Msg("Failed to start http server")
			os.Exit(52)
		}
	}()
	close(httpServerStarted)

	stop := <-gracefulStop
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Logger.Error().Err(err).Msg("Failed to shutdown http server")
	}
	logger.Logger.Debug().Msgf("http server stopped. Received signal %s", stop)
	close(httpServerStopped)
}
