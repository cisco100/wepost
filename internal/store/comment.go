package store

import (
	"context"
	"database/sql"
)

type CommentStore struct {
	db *sql.DB
}

func (cs *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `INSERT INTO comments (id,post_id,user_id,comment) VALUES($1,$2,$3,$4) RETURNING created_at`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := cs.db.QueryRowContext(
		ctx,
		query,
		comment.ID,
		comment.PostID,
		comment.UserID,
		comment.Comment,
	).Scan(&comment.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CommentStore) GetPostWithComment(ctx context.Context, postID string) ([]Comment, error) {
	query := `SELECT comments.id,comments.post_id,comments.user_id,comments.comment,users.id,users.username FROM comments 
	JOIN  users ON users.id=comments.user_id 
	WHERE comments.post_id=$1 
	ORDER BY comments.created_at DESC`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := cs.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var comment Comment
		comment.User = User{}
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Comment,
			&comment.User.ID,
			&comment.User.Username,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil

}
