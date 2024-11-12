package postgresrep

import (
	"context"
	"jobsearcher_user/pkg/client"
)

type PgxManager struct {
	db   client.Client
	ctxm client.CtxManager
}

func newPgxManager(db client.Client) client.DBManager {
	return &PgxManager{
		db:   db,
		ctxm: client.NewPgxCtxManager(db),
	}
}
func (pm PgxManager) DefaultTrOrDB(ctx context.Context) (client.Client, error) {
	ctxClient := pm.ctxm.Default(ctx)
	if ctxClient != nil {
		return ctxClient, nil
	}
	return pm.db, nil
}

func (pm PgxManager) TrOrDB(ctx context.Context, key interface{}) (client.Client, error) {
	ctxClient := pm.ctxm.ByKey(ctx, key)
	if ctxClient != nil {
		return ctxClient, nil
	}
	return pm.db, nil
}
