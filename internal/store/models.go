package store

type User struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  Password `json:"-"`
	CreatedAt string   `json:"created_at"`
	IsActive  bool     `json:"is_active"`
	Role      Role     `json:"role"`
}

type Password struct {
	PlainText *string
	Hash      []byte
}
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	UserID    string    `json:"user_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Version   int       `json:"version"`
	Comment   []Comment `json:"coment"`
	User      User      `json:"user"`
}

type Comment struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type Follower struct {
	UserID     string `json:"user_id"`
	FollowerID string `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

type PostWithMetaData struct {
	Post
	CountComment int64 `json:"count_comment"`
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}
