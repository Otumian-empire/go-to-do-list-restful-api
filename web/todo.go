package web

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
)

type TodoController struct {
	model model.TodoModel
}

func (controller *TodoController) CreateTodo() gin.HandlerFunc {
	return func(context *gin.Context) {
		value, isValue := context.MustGet("user").(model.User)

		log.Println(value)

		if !isValue {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		var payload CreateTodoRequestBody

		if err := context.BindJSON(&payload); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		payload.Task = strings.Trim(payload.Task, " ")

		if len(payload.Task) < 1 {
			context.JSON(FailureMessageResponse(INVALID_TODO))
			return
		}

		if err := controller.model.CreateTodo(value.Id, payload.Task); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(SuccessMessageResponse(TODO_CREATED_SUCCESSFULLY))
		return
	}
}

// func (controller *TodoController) ReadTodo() gin.HandlerFunc

// func (controller *TodoController) ReadTodos() gin.HandlerFunc

// func (controller *TodoController) UpdateTodoTask() gin.HandlerFunc

// func (controller *TodoController) UpdateTodoCompleted() gin.HandlerFunc

// func (controller *TodoController) DeleteTodo() gin.HandlerFunc
