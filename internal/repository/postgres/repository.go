package postgresrep

import (
	ipsql "jobsearcher_user/internal/app/repository/psql"
	"jobsearcher_user/pkg/client"
)

type Repository struct {
	TxManager client.Manager
	Keyword   ipsql.Keyword
	Vacancy   ipsql.Vacancy
}

func New(pg client.Client) Repository {
	return Repository{
		TxManager: NewPgxTxManager(pg),
		Keyword:   NewKeyword(pg),
		Vacancy:   NewVacancy(pg),
	}
}
