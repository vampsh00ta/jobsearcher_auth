package psql

import (
	"context"
	"jobsearcher_user/internal/entity"
)

type Keyword interface {
	GetAll(ctx context.Context) ([]entity.Keyword, error)
}
type Vacancy interface {
	AddKeywords(ctx context.Context, jobSlugID string, slugNames []string) error
}
