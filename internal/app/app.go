package app

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"jobsearcher_auth/config"
	"jobsearcher_auth/internal/http"
	psqlrep "jobsearcher_auth/internal/repository/postgres"
	"jobsearcher_auth/internal/service"
	"net"

	// psqlrep "jobsearcher/internal/repository/postgres_repository"
	// "jobsearcher/internal/service".

	"jobsearcher_auth/pkg/client"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = logger.Sync()
	}()
	return logger
}
func Run(cfg *config.Config) {
	ctx := context.Background()
	defer ctx.Done()

	logger := NewLogger()
	pg, err := client.NewPostgresClient(ctx, 5, cfg.PG)
	if err != nil {
		logger.Fatal("jobsearcher - Run - postgres.New: %v", zap.Error(err))
	}
	defer pg.Close()

	// Psql Repository
	logger.Info("starting repository...")
	psqlRep := psqlrep.New(pg)

	// Services
	logger.Info("starting services...")
	srvc := service.New(psqlRep, cfg)

	lis, err := net.Listen("tcp", ":"+cfg.HTTP.Port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	authServer := http.NewAuthGRPC(srvc.Auth, srvc.Link)
	http.Register(s, authServer)
	if err = s.Serve(lis); err != nil {
		logger.Fatal("Ошибка запуска сервера: %v", zap.Error(err))
	}

}
