package transaction

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrTransaction = errors.New("transaction error")

type ContextKey struct{}

type Manager struct {
	pool *pgxpool.Pool
}

func NewManager(pool *pgxpool.Pool) *Manager {
	return &Manager{pool: pool}
}

func (tm *Manager) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	if tx := GetTransaction(ctx); tx != nil {
		return fn(ctx)
	}

	tx, err := tm.pool.Begin(ctx)
	if err != nil {
		return errors.Join(ErrTransaction, err)
	}

	txCtx := context.WithValue(ctx, ContextKey{}, tx)

	err = fn(txCtx)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return errors.Join(ErrTransaction, rbErr)
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Join(ErrTransaction, err)
	}

	return nil
}

func (tm *Manager) GetQueryDoer(ctx context.Context) interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
} {
	if tx := GetTransaction(ctx); tx != nil {
		return tx
	}
	return tm.pool
}

// GetTransaction retrieves the transaction from context
func GetTransaction(ctx context.Context) pgx.Tx {
	tx, ok := ctx.Value(ContextKey{}).(pgx.Tx)
	if !ok {
		return nil
	}
	return tx
}
