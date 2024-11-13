package postgresrep

import (
	ipsql "jobsearcher_user/internal/app/repository/psql"
	"jobsearcher_user/pkg/client"
)

type Repository struct {
	TxManager client.Manager
	Link      ipsql.Link
}

func New(pg client.Client) Repository {
	return Repository{
		TxManager: NewPgxTxManager(pg),
		Link:      NewLink(pg),
	}
}
