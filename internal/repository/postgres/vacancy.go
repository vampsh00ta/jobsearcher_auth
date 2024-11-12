package postgresrep

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	ipsql "jobsearcher_user/internal/app/repository/psql"
	"jobsearcher_user/pkg/client"
)

type vacancy struct {
	db client.DBManager
}

func NewVacancy(c client.Client) ipsql.Vacancy {
	return &vacancy{
		db: newPgxManager(c),
	}
}

func (j vacancy) getVacancyIDBySlugID(ctx context.Context, jobSlug string) (int, error) {
	var err error
	client, err := j.db.DefaultTrOrDB(ctx)
	if err != nil {
		return -1, err
	}

	q := `select id from vacancy where slug_id = $1`

	var id int
	if err = client.QueryRow(ctx, q, jobSlug).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
func (j vacancy) AddKeywords(ctx context.Context, vacancySlugID string, keywords []string) error {
	var err error

	client, err := j.db.DefaultTrOrDB(ctx)
	if err != nil {
		return err
	}
	var q string
	if len(keywords) == 0 || keywords == nil {
		return nil
	}
	jobID, err := j.getVacancyIDBySlugID(ctx, vacancySlugID)
	if err != nil {
		return err
	}
	q = `select id from keyword where name  = any($1)`
	inputVals := make([]any, 0, len(keywords))

	rows, err := client.Query(ctx, q, keywords)
	if err != nil {
		return err
	}
	keywordIDs, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return err
	}
	if len(keywordIDs) == 0 {
		return nil
	}
	q = `insert into  vacancy_keyword (vacancy_id,keyword_id) values 
		 
		 `
	inputVals = make([]any, 0, 1+len(keywordIDs))
	inputVals = append(inputVals, jobID)

	for i, slugID := range keywordIDs {
		q += fmt.Sprintf("($%d,$%d),", 1, i+2)
		inputVals = append(inputVals, slugID)
	}
	q = q[:len(q)-1]
	q += ` on conflict do nothing`
	if err = client.QueryRow(ctx, q, inputVals...).Scan(); dbError(err) != nil {
		return err
	}
	return nil
}
