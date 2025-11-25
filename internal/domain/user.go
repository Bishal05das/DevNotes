package domain

type User struct {
	ID         int    `db:"id" json:"user_id"`
	Name       string `db:"name" json:"name"`
	IsAdmin    bool   `db:"is_admin" json:"is_admin"`
	Reputation int    `db:"reputation" json:"reputation"`
}

type UserRepository interface {
    CreateUser(user *User) error
    GetUser(id int) (*User, error)
    ListUser() ([]*User, error)
    DeleteUser(id int) error
    UpdateReputation(id int) error
    UpdateUser(u *User) (*User,error)
}