package model

import "github.com/otumian-empire/go-to-do-list-restful-api/config"

type Todo struct {
	Id        config.IdType `db:"id" json:"id"`
	Task      string        `db:"task" json:"task"`
	Completed bool          `db:"completed" json:"completed"`
	CreatedAt string        `db:"created_at" json:"created_at"`
	UpdatedAt string        `db:"updated_at" json:"updated_at"`
	User      config.IdType `db:"user" json:"user"`
}

type TodoModel interface {
	CreateTodo(userId config.IdType, task string) error
	ReadTodoById(id config.IdType) (Todo, error)
	PaginateTodo(userId config.IdType, offset, limit int) ([]Todo, error)
	UpdateTodoTask(userId, id config.IdType, task string) error
	UpdateTodoCompletedState(userId, id config.IdType, completed bool) error
	DeleteTodo(userId, id config.IdType) error
	DeleteTodos(userId config.IdType) error
}
