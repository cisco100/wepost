package store

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
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
		user.Password.Hash,
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

func (us *UserStore) ActivateAccount(ctx context.Context, token string, invite_expiry time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	WithTrxn(us.db, ctx, func(tx *sql.Tx) error {
		user, err := us.getUserFromInvite(ctx, tx, token, invite_expiry)
		if err != nil {
			return err
		}

		user.IsActive = true
		if err := us.updateUser(ctx, tx, user); err != nil {
			return err
		}

		if err := us.cleanInvite(ctx, tx, user.ID); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (us *UserStore) cleanInvite(ctx context.Context, tx *sql.Tx, userID string) error {
	query := `DELETE FROM user_invitations WHERE user_id=$1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return err
	}
	return nil

}

func (us *UserStore) updateUser(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `UPDATE users SET username=$1,email=$2,is_active=$3
	WHERE id=$4
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := tx.ExecContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.IsActive,
		user.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
func (us *UserStore) getUserFromInvite(ctx context.Context, tx *sql.Tx, token string, invite_expiry time.Time) (*User, error) {
	hash := sha256.Sum256([]byte(token))
	hashToken := hex.EncodeToString(hash[:])

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `SELECT users.id,users.username,users.email,users.created_at,users.is_active FROM users 
	JOIN user_invitations ON users.id=user_invitations.user_id
	WHERE user_invitations.token=$1 AND user_invitations.invite_expiry>$2
	`
	user := &User{}
	err := tx.QueryRowContext(
		ctx,
		query,
		hashToken,
		invite_expiry,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.IsActive,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return user, nil

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
