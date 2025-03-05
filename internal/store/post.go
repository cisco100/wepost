package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

func (ps *PostStore) Create(ctx context.Context, post *Post) error {

	query := `INSERT INTO posts(id,title,content,tags,user_id) VALUES($1,$2,$3,$4,$5) RETURNING id,user_id,created_at,updated_at,version`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := ps.db.QueryRowContext(
		ctx,
		query,
		post.ID,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.UserID,
	).Scan(
		&post.ID,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
	)

	if err != nil {
		return err
	}
	return nil

}

func (ps *PostStore) GetPostById(ctx context.Context, postID string) (*Post, error) {
	var post Post
	query := `SELECT * FROM posts WHERE id=$1`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	err := ps.db.QueryRowContext(
		ctx,
		query,
		postID,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &post, nil

}

func (ps *PostStore) AllPost(ctx context.Context) ([]Post, error) {
	query := `SELECT * FROM posts`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := ps.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			pq.Array(&post.Tags),
			&post.UserID,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Version,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)

	}

	return posts, nil
}

func (ps *PostStore) DeletePost(ctx context.Context, postID string) error {

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	query := `DELETE FROM posts WHERE id=$1`

	res, err := ps.db.ExecContext(
		ctx,
		query,
		postID,
	)

	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (ps *PostStore) UpdatePost(ctx context.Context, post *Post) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	query := `
	UPDATE posts
	SET title = $1, content = $2, tags=$3,updated_at=CURRENT_TIMESTAMP,version=version+1	WHERE id = $4 and version=$5
	RETURNING id,version
	`
	err := ps.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.ID,
		post.Version,
	).Scan(
		&post.ID,
		&post.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}

	}
	return nil
}

func (ps *PostStore) GetUserFeed(ctx context.Context, userID string, pag PaginatedFeedQuery) ([]PostWithMetaData, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	// 	query := `SELECT posts.id, posts.title, posts.content, posts.created_at, posts.version, posts.tags,
	//        users.username, COUNT(comments.id) AS count_comment
	// FROM posts
	// LEFT JOIN users ON posts.user_id = users.id
	// LEFT JOIN comments ON comments.post_id = posts.id
	// LEFT JOIN followers ON followers.follower_id = posts.user_id
	// WHERE followers.user_id = $1
	//       OR posts.user_id = $1
	// GROUP BY posts.id, users.username
	// ORDER BY posts.created_at ` + pag.Sort + `
	// LIMIT $2 OFFSET $3;
	// `

	query := `
		SELECT 
			posts.id, posts.user_id, posts.title, posts.content, posts.created_at, posts.version, posts.tags,
			users.username,
			COUNT(comments.id) AS count_comment
		FROM posts 
		LEFT JOIN comments  ON comments.post_id = posts.id
		LEFT JOIN users  ON posts.user_id = users.id
		JOIN followers  ON followers.follower_id = posts.user_id OR posts.user_id = $1
		WHERE 
			followers.user_id = $1 AND
			(posts.title ILIKE '%' || $4 || '%' OR posts.content ILIKE '%' || $4 || '%') AND
			(posts.tags @> $5 OR $5 = '{}')
		GROUP BY posts.id, users.username
		ORDER BY posts.created_at ` + pag.Sort + `
		LIMIT $2 OFFSET $3
	`
	rows, err := ps.db.QueryContext(ctx, query, userID, pag.Limit, pag.Offset, pag.Search, pq.Array(pag.Tags))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	feeds := []PostWithMetaData{}
	for rows.Next() {
		var feed PostWithMetaData
		err := rows.Scan(
			&feed.Post.ID,
			&feed.Post.Title,
			&feed.Post.Content,
			&feed.Post.CreatedAt,
			&feed.Post.Version,
			pq.Array(&feed.Post.Tags),
			&feed.User.Username,
			&feed.CountComment,
		)
		if err != nil {
			return nil, err
		}

		feeds = append(feeds, feed)
	}
	return feeds, nil

}
