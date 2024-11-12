package postgresrep

import (
	"context"
	"github.com/jackc/pgx/v5"
	ipsql "jobsearcher_user/internal/app/repository/psql"
	"jobsearcher_user/internal/entity"
	"jobsearcher_user/pkg/client"
)

type keyword struct {
	db client.DBManager
}

func NewKeyword(c client.Client) ipsql.Keyword {
	return &keyword{
		db: newPgxManager(c),
	}
}
func (s keyword) GetAll(ctx context.Context) ([]entity.Keyword, error) {
	client, err := s.db.DefaultTrOrDB(ctx)
	if err != nil {
		return nil, err
	}

	q := `
		select id,name,slug from keyword
		
	`

	rows, err := client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rowModels, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Keyword])
	if err != nil {
		return nil, err
	}
	return rowModels, nil
}
