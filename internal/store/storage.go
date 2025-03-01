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
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Post interface {
		Create(context.Context, *Post) error
		GetPostById(context.Context, string) (*Post, error)
		AllPost(context.Context) ([]Post, error)
		UpdatePost(context.Context, *Post) error
		DeletePost(context.Context, string) error
	}
	User interface {
		Create(context.Context, *User) error
		GetUserById(context.Context, string) (*User, error)
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
