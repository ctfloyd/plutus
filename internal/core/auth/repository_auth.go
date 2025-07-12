package user

import (
	"context"
	"plutus/internal/common/transaction"
	"plutus/internal/gen/db"
)

type Repository interface {
	InsertAuth(ctx context.Context, params db.CreateAuthParams) error
	InsertToken(ctx context.Context, params db.CreateTokenParams) error
	GetAuth(ctx context.Context, userId string) (db.Auth, error)
	GetToken(ctx context.Context, token string) (db.Token, error)
	RevokeToken(ctx context.Context, token string) error
}

type RdsRepository struct {
	queries *db.Queries
	txMgr   *transaction.Manager
}

func NewRdsRepository(queries *db.Queries, txMgr *transaction.Manager) *RdsRepository {
	return &RdsRepository{queries: queries, txMgr: txMgr}
}

func (r *RdsRepository) GetToken(ctx context.Context, token string) (db.Token, error) {
	return r.getQueries(ctx).GetToken(ctx, token)
}

func (r *RdsRepository) RevokeToken(ctx context.Context, token string) error {
	return r.getQueries(ctx).RevokeToken(ctx, token)
}

func (r *RdsRepository) GetAuth(ctx context.Context, userId string) (db.Auth, error) {
	return r.getQueries(ctx).GetAuth(ctx, userId)
}

func (r *RdsRepository) InsertAuth(ctx context.Context, params db.CreateAuthParams) error {
	_, err := r.getQueries(ctx).CreateAuth(ctx, params)
	return err
}

func (r *RdsRepository) InsertToken(ctx context.Context, params db.CreateTokenParams) error {
	_, err := r.getQueries(ctx).CreateToken(ctx, params)
	return err
}

func (r *RdsRepository) getQueries(ctx context.Context) *db.Queries {
	if tx := transaction.GetTransaction(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}

	return r.queries
}
