package main

import (
	"context"
	apiGRPC "github.com/krls256/card-validator/api/grpc"
	"github.com/krls256/card-validator/pkg/config"
	"github.com/krls256/card-validator/pkg/handlers"
	"github.com/krls256/card-validator/pkg/transport/grpc"
	"github.com/krls256/card-validator/pkg/transport/http"
	"github.com/krls256/card-validator/utils"
	"log"
	"log/slog"
	"os"
	"time"
)

func main() {
	now := time.Now()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger = logger.With("service_name", "card_validator")

	slog.SetDefault(logger)

	cfg, err := config.New("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer[apiGRPC.CardValidatorServiceServer](
		cfg.GRPCConfig, handlers.NewGRPCCardValidatorHandler(),
		apiGRPC.RegisterCardValidatorServiceServer)

	grpcServer.RunAsync()

	httpServer := http.NewServer(context.Background(), "card_validator", slog.Default(),
		cfg.HTTPConfig, []http.Handler{handlers.NewCardHTTPValidatorHandler()})

	httpServer.AsyncRun()

	<-utils.WaitTermSignal()

	slog.Info("shutdown", slog.Duration("server was running for", time.Since(now)))
}
