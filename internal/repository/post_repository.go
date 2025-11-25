package repository

import (
	"database/sql"

	"github.com/bishal05das/blog_app_1/internal/domain"
	"github.com/jmoiron/sqlx"
)

type PostRepositoryDB struct {
	db *sqlx.DB
}

func NewPostRepositoryDB(db *sqlx.DB) *PostRepositoryDB {
	return &PostRepositoryDB{
		db: db,
	}
}

func (h *PostRepositoryDB) CreatePost(post *domain.Post) error {
	query := `INSERT INTO posts (title,user_id) VALUES ($1,$2) RETURNING id;`

	return h.db.QueryRow(query, post.Title, post.UserID).Scan(&post.ID)
}

func (h *PostRepositoryDB) GetPost(id int) (*domain.Post, error) {
	var post domain.Post
	query := `SELECT id,title,user_id,score FROM posts WHERE id=$1;`
	err := h.db.Get(&post, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (h *PostRepositoryDB) ListPost() ([]*domain.Post, error) {
	var posts []*domain.Post
	query := `SELECT id,title,user_id,score FROM posts;`
	err := h.db.Select(&posts, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return posts, nil

}

func (h *PostRepositoryDB) DeletePost(id int) error {
	query := `DELETE FROM posts WHERE id=$1`
	_,err := h.db.Exec(query,id)
	if err != nil {
		return err
	}
	return nil

}

func (h *PostRepositoryDB) UpdatePost(p domain.Post) (*domain.Post,error) {
	query := `Update posts SET title=$1,user_id=$2 score=$3 WHERE id=$4;`
	row := h.db.QueryRow(query,p.Title,p.UserID,p.Score,p.ID)
	err := row.Err()
	if err != nil {
		return nil,err
	}
	return &p, nil
} 

func (h *PostRepositoryDB) UpdateScore(id int) error{
	query := `Upadte posts SET score = score + 1 WHERE id=$1; `
	row := h.db.QueryRow(query,id)
	err := row.Err()

	if err != nil {
		return err
	}
	return nil
}

func (h *PostRepositoryDB) GetAuthorIDByPostID(id int) (*int,error) {
	var userID int
	query := `SELECT user_id FROM posts WHERE id=$1;`
	err := h.db.QueryRow(query, id).Scan(&userID)
	if err != nil {
		return nil,err
	}
	return &userID,nil
}