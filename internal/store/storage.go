package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	ErrDuplicateEmail    = errors.New("a user with the same email already exists")
	ErrDuplicateUsername = errors.New("a usser with the same username already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Post interface {
		Create(context.Context, *Post) error
		GetPostById(context.Context, string) (*Post, error)
		AllPost(context.Context) ([]Post, error)
		UpdatePost(context.Context, *Post) error
		DeletePost(context.Context, string) error
		GetUserFeed(context.Context, string, PaginatedFeedQuery) ([]PostWithMetaData, error)
	}
	User interface {
		create(context.Context, *sql.Tx, *User) error
		updateUser(context.Context, *sql.Tx, *User) error
		cleanInvite(context.Context, *sql.Tx, string) error
		getUserFromInvite(context.Context, *sql.Tx, string, time.Time) (*User, error)
		CreateAndInvite(context.Context, *User, string, time.Duration) error
		GetUserById(context.Context, string) (*User, error)
		ActivateAccount(context.Context, string, time.Time) error
		DeleteUser(context.Context, string) error
	}

	Comment interface {
		Create(context.Context, *Comment) error
		GetPostWithComment(context.Context, string) ([]Comment, error)
	}
	Follower interface {
		Follow(context.Context, string, string) error
		Unfollow(context.Context, string, string) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Post:     &PostStore{db},
		User:     &UserStore{db},
		Comment:  &CommentStore{db},
		Follower: &FollowStore{db},
	}
}
