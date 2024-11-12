package service

import (
	"context"
	"jobsearcher_user/internal/entity"
)

type Vacancy interface {
	AddKeywords(ctx context.Context, v *entity.Vacancy) error
}
type Auth interface {
	VerifyToken(_ context.Context, accessToken string) (bool, error)
	CreateToken(ctx context.Context, user entity.User) (string, error)
}
