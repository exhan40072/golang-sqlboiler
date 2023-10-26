package mysql

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type contextKey string

const transactionKey contextKey = "transaction"

// DB ...
type DB struct {
	db *sql.DB
}

// NewDB ...
func NewDB() *DB {
	return &DB{getDBInstance()}
}

// GetConnection ...
func (r *DB) GetConnection(ctx context.Context) boil.ContextExecutor {
	tx, ok := getTxFromContext(ctx)
	if !ok {
		return r.db
	}
	return tx
}

// getTxFromContext ...
func getTxFromContext(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(transactionKey).(*sql.Tx)
	return tx, ok
}

// Transaction ...
func (r *DB) Transaction(ctx context.Context, f func(context.Context) error) (err error) {
	tx, err := boil.BeginTx(ctx, nil) // TODO: 第二引数についてもう少し調べる
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	ctx = context.WithValue(ctx, transactionKey, tx)
	if err = f(ctx); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
