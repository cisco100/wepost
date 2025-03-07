package store

import (
	"context"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users(id,username,email,password) VALUES($1,$2,$3,$4) RETURNING id,created_at `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	err := us.db.QueryRowContext(
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
		return err
	}

	return nil
}

func (us *UserStore) CreateAndInvite(ctx context.Context, user *User, token string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
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
