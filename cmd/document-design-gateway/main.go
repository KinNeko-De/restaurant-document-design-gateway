package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/github/oauth"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation/health"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/operation/logger"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/router"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	googleHealth "google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
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

	httpServerStarted := make(chan struct{})
	httpServerStopped := make(chan struct{})
	grpcServerStarted := make(chan struct{})
	grpcServerStopped := make(chan struct{})
	go startHttpServer(httpServerStarted, httpServerStopped, ":8080")
	go startGrpcServer(grpcServerStarted, grpcServerStopped, ":3110")

	<-grpcServerStarted
	<-httpServerStarted
	health.Ready()

	<-grpcServerStopped
	<-httpServerStopped
	logger.Logger.Info().Msg("Application stopped.")
	os.Exit(0)
}

func startHttpServer(httpServerStarted chan struct{}, httpServerStopped chan struct{}, port string) {
	router := router.SetupRouter()
	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)
	logger.Logger.Debug().Msg("starting http server")

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
	close(httpServerStopped)
}

func startGrpcServer(grpcServerStarted chan struct{}, grpcServerStopped chan struct{}, port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Logger.Error().Err(err).Msgf("Failed to listen on port %v", port)
		os.Exit(51)
	}

	grpcServer := configureGrpcServer()
	healthServer := googleHealth.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	health.Initialize(healthServer)

	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)
	logger.Logger.Debug().Msg("starting grpc server")

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			logger.Logger.Error().Err(err).Msg("failed to serve grpc server")
			os.Exit(52)
		}
	}()
	close(grpcServerStarted)

	stop := <-gracefulStop
	healthServer.Shutdown()
	grpcServer.GracefulStop()
	logger.Logger.Debug().Msgf("http server stopped. received signal %s", stop)
	close(grpcServerStopped)
}

func configureGrpcServer() *grpc.Server {
	// Handling of panic to prevent crash from example nil pointer exceptions
	logPanic := func(p any) (err error) {
		log.Error().Any("method", p).Err(err).Msg("Recovered from panic.")
		return status.Errorf(codes.Internal, "Internal server error occured.")
	}

	opts := []recovery.Option{
		recovery.WithRecoveryHandler(logPanic),
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			recovery.UnaryServerInterceptor(opts...),
		),
		grpc.StreamInterceptor(
			recovery.StreamServerInterceptor(opts...),
		),
	)
	RegisterAllGrpcServices(grpcServer)
	return grpcServer
}

func RegisterAllGrpcServices(grpcServer *grpc.Server) {
}
