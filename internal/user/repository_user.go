package user

import (
	"context"
	"github.com/cockroachdb/errors"
	"plutus/internal/common/transaction"
	"plutus/internal/gen/db"
)

type Repository interface {
	IsEmailTaken(ctx context.Context, email string) (bool, error)
	GetUserById(ctx context.Context, id string) (db.User, error)
	InsertUser(ctx context.Context, user db.CreateUserParams) (db.User, error)
	UpdateUser(ctx context.Context, user db.UpdateUserParams) (db.User, error)
}

type RdsRepository struct {
	queries *db.Queries
	txMgr   *transaction.Manager
}

func NewRdsRepository(queries *db.Queries, txMgr *transaction.Manager) *RdsRepository {
	return &RdsRepository{queries: queries, txMgr: txMgr}
}

func (r *RdsRepository) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	taken, err := r.getQueries(ctx).IsEmailTaken(ctx, email)
	return taken, errors.WithStack(err)
}

func (r *RdsRepository) GetUserById(ctx context.Context, id string) (db.User, error) {
	data, err := r.getQueries(ctx).GetUser(ctx, id)
	return data, errors.WithStack(err)
}

func (r *RdsRepository) InsertUser(ctx context.Context, user db.CreateUserParams) (db.User, error) {
	data, err := r.getQueries(ctx).CreateUser(ctx, user)
	return data, errors.WithStack(err)
}

func (r *RdsRepository) UpdateUser(ctx context.Context, user db.UpdateUserParams) (db.User, error) {
	data, err := r.getQueries(ctx).UpdateUser(ctx, user)
	return data, errors.WithStack(err)
}

func (r *RdsRepository) getQueries(ctx context.Context) *db.Queries {
	if tx := transaction.GetTransaction(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}

	return r.queries
}
