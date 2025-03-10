package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) create(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `INSERT INTO users(id,username,email,password) VALUES($1,$2,$3,$4) RETURNING id,created_at `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	err := tx.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case err.Error() == `pq:duplicate key value violates user constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq:duplicate key value violates user constraint "users_username_key"`:
			return ErrDuplicateUsername
		default:
			return err

		}
	}

	return nil
}

func (us *UserStore) CreateAndInvite(ctx context.Context, user *User, token string, invite_expiry time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	WithTrxn(us.db, ctx, func(tx *sql.Tx) error {
		if err := us.create(ctx, tx, user); err != nil {
			return err
		}

		err := us.createUserInvite(ctx, tx, user.ID, token, invite_expiry)
		if err != nil {
			return err
		}
		return nil

	})
	return nil
}

func (us *UserStore) GetUserById(ctx context.Context, userID string) (*User, error) {
	query := `SELECT id,username,email,created_at FROM users WHERE id=$1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	var user User

	err := us.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil

}
func (pass *Password) Set(plainText string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	pass.PlainText = &plainText
	pass.Hash = hash
	return nil
}
