package app

import (
	"context"
	"go.uber.org/zap"
	"jobsearcher_user/config"
	"jobsearcher_user/internal/http"
	psqlrep "jobsearcher_user/internal/repository/postgres"
	"jobsearcher_user/internal/service"

	// psqlrep "jobsearcher/internal/repository/postgres_repository"
	// "jobsearcher/internal/service".

	"jobsearcher_user/pkg/client"
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
	app := newHTTP(cfg, logger)
	http.New(app, http.Services{
		Auth: srvc.Auth,
		Link: srvc.Link,
	})
	//u := entity.User{
	//	ID:        564764193,
	//	FirstName: "name",
	//	LastName:  "",
	//	Username:  "vamp_sh00ta",
	//	PhotoUrl:  "https%3A%2F%2Ft.me%2Fi%2Fuserpic%2F320%2Fg5_nfUkP_Gw0M0P27NFukf34YYQOX0m87CfUVel4CEM.jpg"}
	//access, err := srvc.Auth.CreateToken(ctx, u)
	//fmt.Println(access, err)
	//fmt.Println(srvc.Auth.VerifyToken(ctx, access))
	if err = app.Listen(":" + cfg.HTTP.Port); err != nil {
		logger.Fatal("Ошибка запуска сервера: %v", zap.Error(err))
	}

}
