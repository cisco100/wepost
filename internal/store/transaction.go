package store

import (
	"context"
	"database/sql"
)

func WithTrxn(db *sql.DB, ctx context.Context, fxn func(*sql.Tx) error) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fxn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
