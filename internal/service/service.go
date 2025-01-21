package service

import (
	"jobsearcher_auth/config"
	isrvc "jobsearcher_auth/internal/app/service"
	postgresrep "jobsearcher_auth/internal/repository/postgres"
)

type Service struct {
	Auth isrvc.Auth
	Link isrvc.Link
}

func New(rep postgresrep.Repository, cfg *config.Config) Service {
	return Service{
		Auth: NewAuth(cfg),
		Link: NewLink(cfg, rep.Link, rep.TxManager),
	}
}
