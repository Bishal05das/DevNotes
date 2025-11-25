package repository


import (
	"database/sql"

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
	query := `Update users SET name=$1,is_admin=$2,reputation=$3  WHERE id=$4`
	row := h.db.QueryRow(query,u.Name,u.IsAdmin,u.Reputation,u.ID)
	err := row.Err()
	if err != nil {
		return nil,err
	}
	return u, nil
} 


func (h *UserRepositoryDB) UpdateReputation(id int) error {
	query := `Upadte users SET reputation = reputation + 5 WHERE id=$1; `
	row := h.db.QueryRow(query,id)
	err := row.Err()

	if err != nil {
		return err
	}
	return nil
}