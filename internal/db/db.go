package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/cisco100/wepost/internal/store"
)

func NewConnection(addr string, maxOpenCon int, maxIdleCon int, maxIdleTime string) (*sql.DB, error) {
	duration, _ := time.ParseDuration(maxIdleTime)

	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenCon)
	db.SetMaxIdleConns(maxIdleCon)
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), store.QueryTimeoutDuration)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, err
}
