package store

import (
	"context"
	"database/sql"
	"log"
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
		log.Println("Transaction rolled back:", err) // Debugging
		return err
	}
	log.Println("Transaction committed successfully") // Debugging
	return tx.Commit()
}
