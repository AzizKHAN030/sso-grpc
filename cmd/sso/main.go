package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/azizkhan030/sso-grpc/internal/app"
	"github.com/azizkhan030/sso-grpc/internal/config"
	"github.com/azizkhan030/sso-grpc/internal/lib/logger/handlers/slogpretty"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()

	log := setupLogger(config.Env)

	log.Info("starting application")

	application := app.New(log, config.GRPC.Port, config.StoragePath, config.TokenTTL)

	go application.GRPCServer.MustRun()

	// TODO app initialization
	// TODO start gRPC server

	//Graceful shutDown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	receivedSignal := <-stop

	log.Info("stopping application", slog.String("signal", receivedSignal.String()))

	application.GRPCServer.Stop()
	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}
