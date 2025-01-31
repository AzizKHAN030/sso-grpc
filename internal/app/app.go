package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/azizkhan030/sso-grpc/internal/app/grpc"
	"github.com/azizkhan030/sso-grpc/internal/services/auth"
	"github.com/azizkhan030/sso-grpc/internal/storage/sqlite"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)

	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
