package model

import "github.com/otumian-empire/go-to-do-list-restful-api/config"

type Todo struct {
	Id        config.IdType `db:"id" json:"id"`
	Task      string        `db:"task" json:"task"`
	Completed bool          `db:"completed" json:"completed"`
	CreatedAt string        `db:"created_at" json:"created_at"`
	UpdatedAt string        `db:"updated_at" json:"updated_at"`
}

type TodoModel interface {
	CreateTodo(task string) error
	ReadById(id config.IdType) (Todo, error)
	Paginate(offset, limit int) []Todo
	UpdateTask(id config.IdType, task string) error
	UpdateCompletedState(id config.IdType, completed bool) error
	DeleteTodo(id config.IdType) error
}
