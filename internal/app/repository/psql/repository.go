package psql

import (
	"context"
)

type Link interface {
	Create(ctx context.Context, hash, accessToken string) error
	GetAccessToken(ctx context.Context, hash string) (string, error)
	Delete(ctx context.Context, hash string) error
}
