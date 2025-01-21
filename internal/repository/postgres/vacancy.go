package postgresrep

import (
	"context"
	ipsql "jobsearcher_auth/internal/app/repository/psql"
	"jobsearcher_auth/pkg/client"
)

type link struct {
	db client.DBManager
}

func NewLink(c client.Client) ipsql.Link {
	return &link{
		db: newPgxManager(c),
	}
}

func (l link) Create(ctx context.Context, hash, accessToken string) error {
	var err error
	client, err := l.db.DefaultTrOrDB(ctx)
	if err != nil {
		return err
	}

	q := `insert into link(hash,access_token) values ($1,$2) on conflict (hash) do nothing  `

	if err = client.QueryRow(ctx, q, hash, accessToken).Scan(); dbError(err) != nil {
		return err
	}
	return nil
}

func (l link) GetAccessToken(ctx context.Context, hash string) (string, error) {
	var err error
	client, err := l.db.DefaultTrOrDB(ctx)
	if err != nil {
		return "", err
	}

	q := `select access_token from link where hash = $1`

	var acccessToken string
	if err = client.QueryRow(ctx, q, hash).Scan(&acccessToken); err != nil {
		return "", err
	}
	return acccessToken, nil
}

func (l link) Delete(ctx context.Context, hash string) error {
	var err error
	client, err := l.db.DefaultTrOrDB(ctx)
	if err != nil {
		return err
	}

	q := `delete from link where hash = $1`

	if err = client.QueryRow(ctx, q, hash).Scan(); dbError(err) != nil {
		return err
	}
	return nil
}
