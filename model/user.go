package model

import "github.com/otumian-empire/go-to-do-list-restful-api/config"

type User struct {
	Id        config.IdType `db:"id" json:"id"`
	Username  string        `db:"username" json:"username"`
	Password  string        `db:"password" json:"password"`
	CreatedAt string        `db:"created_at" json:"created_at"`
	UpdatedAt string        `db:"updated_at" json:"updated_at"`
}

type UserModel interface {
	CreateUser(username, password string) error
	ReadUserById(id config.IdType) (User, error)
	ReadUserByUsername(username string) (User, error)
	UpdateUserPassword(id config.IdType, password string) error
	UpdateUserUsername(id config.IdType, username string) error
	DeleteUser(id config.IdType) error
}
