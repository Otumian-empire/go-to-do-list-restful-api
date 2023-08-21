package web

import (
	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
)

type TodoController struct {
	model model.TodoModel
}

func (controller *TodoController) CreateTodo() gin.HandlerFunc

func (controller *TodoController) ReadTodo() gin.HandlerFunc

func (controller *TodoController) ReadTodos() gin.HandlerFunc

func (controller *TodoController) UpdateTodoTask() gin.HandlerFunc

func (controller *TodoController) UpdateTodoCompleted() gin.HandlerFunc

func (controller *TodoController) DeleteTodo() gin.HandlerFunc
