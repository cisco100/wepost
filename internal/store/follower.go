package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type FollowStore struct {
	db *sql.DB
}

func (fs *FollowStore) Follow(ctx context.Context, userID, followerID string) error {

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `INSERT INTO followers(user_id,follower_id) VALUES($1,$2)`

	_, err := fs.db.ExecContext(
		ctx,
		query,
		userID,
		followerID,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}
	return nil
}

func (fs *FollowStore) Unfollow(ctx context.Context, userID, followerID string) error {

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `DELETE FROM  followers WHERE user_id=$1 AND follower_id=$2`

	_, err := fs.db.ExecContext(
		ctx,
		query,
		userID,
		followerID,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}
	return nil
}
