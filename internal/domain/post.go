package domain

type Post struct {
	ID     int    `db:"id" json:"id"`
	Title  string `db:"title" json:"title"`
	UserID int    `db:"user_id" json:"user_id"`
	Score  int    `db:"score" json:"score"`
}

type PostRepository interface {
	CreatePost(post *Post) error
	GetPost(id int) (*Post, error)
	ListPost() ([]*Post, error)
	DeletePost(id int) error
	UpdatePost(p Post) (*Post,error)
	UpdateScore(id int) error
	GetAuthorIDByPostID(id int) (*int,error)
}
