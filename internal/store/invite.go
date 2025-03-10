package store

import (
	"context"
	"database/sql"
	"time"
)

func (uis *UserStore) createUserInvite(ctx context.Context, tx *sql.Tx, userID string, token string, invite_expiry time.Duration) error {
	query := `INSERT INTO user_invitations(token,user_id,invite_expiry) VALUES($1,$2,$3)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token, userID, time.Now().Add(invite_expiry))
	if err != nil {
		return err
	}
	return nil

}
