package repository

import (
	"database/sql"
	"fmt"

	"github.com/bishal05das/blog_app_1/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) *UserRepositoryDB {
	return &UserRepositoryDB{
		db: db,
	}
}

func (h *UserRepositoryDB) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (name,is_admin) VALUES ($1,$2) RETURNING id;`

	return h.db.QueryRow(query, user.Name,user.IsAdmin).Scan(&user.ID)
}

func (h *UserRepositoryDB) GetUser(id int) (*domain.User, error) {
	var user domain.User
	query := `SELECT id,name,is_admin,reputation FROM users WHERE id=$1`
	err := h.db.Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (h *UserRepositoryDB) ListUser() ([]*domain.User, error) {
	var users []*domain.User
	query := `SELECT id,name,is_admin,reputation FROM users`
	err := h.db.Select(&users, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return users, nil

}

func (h *UserRepositoryDB) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_,err := h.db.Exec(query,id)
	if err != nil {
		return err
	}
	return nil

}

func (h *UserRepositoryDB) UpdateUser(u *domain.User) (*domain.User,error) {
	query := `UPDATE users SET name=$1,is_admin=$2,reputation=$3  WHERE id=$4`
	row := h.db.QueryRow(query,u.Name,u.IsAdmin,u.Reputation,u.ID)
	err := row.Err()
	if err != nil {
		return nil,err
	}
	return u, nil
} 


func (h *UserRepositoryDB) UpdateReputation(id int) error {
	query := `UPDATE users SET reputation = reputation + 5 WHERE id=$1; `
	result,err := h.db.Exec(query,id)
	if err != nil {
		return fmt.Errorf("failed to update reputation: %w",err)
	}

	//Check if any rows were actually updated
	rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("no user found with id %d", id)
    }
    
    //fmt.Printf("Successfully updated reputation for user %d\n", id)
    return nil
}