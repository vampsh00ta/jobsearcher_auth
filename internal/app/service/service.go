package service

import (
	"context"
	"jobsearcher_auth/internal/entity"
)

type Auth interface {
	VerifyToken(_ context.Context, accessToken string) (bool, error)
	CreateToken(ctx context.Context, user entity.User) (string, error)
}

type Link interface {
	Create(ctx context.Context, username, acceessToken string) (string, error)
	Claim(ctx context.Context, hash string) (string, error)
}
