package client

import "context"

type DBManager interface {
	DefaultTrOrDB(ctx context.Context) (Client, error)
	TrOrDB(ctx context.Context, key interface{}) (Client, error)
}

type TxFunc func(ctx context.Context) error
type Manager interface {
	Do(ctx context.Context, f TxFunc) error
}
