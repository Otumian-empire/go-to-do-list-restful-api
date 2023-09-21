package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/otumian-empire/go-to-do-list-restful-api/config"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
)

type Todo struct {
	*sql.DB
}

func (todo *Todo) CreateTodo(userId config.IdType, task string) error {
	result, err := todo.Exec(CREATE_TODO_QUERY, task, false, userId)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf(CREATE_TODO_ERROR)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(CREATE_TODO_ERROR)
	} else if rowsAffected < 1 {
		log.Println(NO_ROW_AFFECT_ERROR)
		return fmt.Errorf(CREATE_TODO_ERROR)
	}

	if _, err := result.LastInsertId(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(CREATE_TODO_ERROR)
	}

	return nil
}

func (todo *Todo) ReadTodoById(userId, id config.IdType) (model.Todo, error) {
	row := todo.QueryRow(READ_TODO_BY_ID_QUERY, userId, id)

	var _todo model.Todo
	err := row.Scan(
		&_todo.Id,
		&_todo.Task,
		&_todo.Completed,
		&_todo.CreatedAt,
		&_todo.UpdatedAt,
		&_todo.User)

	if err != nil {
		log.Println(err)
		return model.Todo{}, fmt.Errorf(NO_ROW_FOUND)
	}

	return _todo, nil
}

func (todo *Todo) PaginateTodo(userId config.IdType, offset, limit int) ([]model.Todo, error) {
	var todos []model.Todo

	rows, err := todo.Query(PAGINATE_TODO_QUERY, userId, offset, limit)

	if err != nil || rows.Err() != nil {
		log.Println(err)
		return nil, fmt.Errorf(PAGINATE_TODO_ERROR)
	}

	defer rows.Close()

	for rows.Next() {
		var _todo model.Todo
		err := rows.Scan(
			&_todo.Id,
			&_todo.Task,
			&_todo.Completed,
			&_todo.CreatedAt,
			&_todo.UpdatedAt,
			&_todo.User)

		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf(PAGINATE_TODO_ERROR)
		}

		todos = append(todos, _todo)
	}

	return todos, nil
}

func (todo *Todo) UpdateTodoTask(userId, id config.IdType, task string) error {
	result, err := todo.Exec(UPDATE_TODO_TASK_QUERY, task, userId, id)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf(UPDATE_TODO_ERROR)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(UPDATE_TODO_ERROR)
	} else if rowsAffected < 1 {
		log.Println(NO_ROW_AFFECT_ERROR)
		return fmt.Errorf(UPDATE_TODO_ERROR)
	}

	return nil
}

func (todo *Todo) UpdateTodoCompletedState(userId, id config.IdType, completed bool) error {
	result, err := todo.Exec(UPDATE_TODO_COMPLETED_STATE_QUERY, completed, userId, id)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf(UPDATE_TODO_ERROR)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(UPDATE_TODO_ERROR)
	} else if rowsAffected < 1 {
		log.Println(NO_ROW_AFFECT_ERROR)
		return fmt.Errorf(UPDATE_TODO_ERROR)
	}

	return nil
}

func (todo *Todo) DeleteTodo(userId, id config.IdType) error {
	result, err := todo.Exec(DELETE_TODO_QUERY, userId, id)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf(DELETE_TODO_ERROR)
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(DELETE_TODO_ERROR)
	} else if rowsAffected < 1 {
		log.Println(NO_ROW_AFFECT_ERROR)
		return fmt.Errorf(DELETE_TODO_ERROR)
	}

	return nil
}

func (todo *Todo) DeleteTodos(userId config.IdType) error {
	result, err := todo.Exec(DELETE_TODOS_QUERY, userId)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf(DELETE_TODO_ERROR)
	}

	if _, err := result.RowsAffected(); err != nil {
		log.Println(err.Error())
		return fmt.Errorf(DELETE_TODO_ERROR)
	}

	// else if rowsAffected < 1 {
	// 	log.Println(NO_ROW_AFFECT_ERROR)
	// 	return fmt.Errorf(DELETE_TODO_ERROR)
	// }

	return nil
}
