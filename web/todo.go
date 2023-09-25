package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
	"github.com/otumian-empire/go-to-do-list-restful-api/repository"
)

type TodoController struct {
	model repository.Repository
}

func (controller *TodoController) CreateTodo() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		log.Println("user", user)
		log.Println("isUser", isUser)
		if !isUser {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		var payload CreateTodoRequestBody

		if err := context.BindJSON(&payload); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		validatedString, isValid := ValidateString(payload.Task)

		if !isValid {
			context.JSON(FailureMessageResponse(INVALID_TODO))
			return
		}

		if err := controller.model.CreateTodo(user.Id, validatedString); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(SuccessMessageResponse(TODO_CREATED_SUCCESSFULLY))
		return
	}
}

func (controller *TodoController) ReadTodo() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		if !isUser {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		var todoId, idErr = ConvertStringIdToInt(context.Param("id"))

		if idErr != nil {
			log.Println(idErr)
			context.JSON(FailureMessageResponse(INVALID_ID))
			return
		}

		todo, todoErr := controller.model.ReadTodoById(user.Id, todoId)

		if todoErr != nil {
			log.Println(todoErr)
			context.JSON(FailureMessageResponse(todoErr.Error()))
			return
		}

		context.JSON(SuccessResponse(TODO_READ_SUCCESSFULLY, todo))
	}
}

func (controller *TodoController) ReadTodos() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		if !isUser {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		pageNumber := convertStringQueryToInt(
			context.DefaultQuery("pageNumber", fmt.Sprintf("%v", DEFAULT_PAGE_NUMBER)),
			DEFAULT_PAGE_NUMBER)

		pageSize := convertStringQueryToInt(
			context.DefaultQuery("pageSize", fmt.Sprintf("%v", DEFAULT_PAGE_SIZE)),
			DEFAULT_PAGE_SIZE)

		var pagination = CleanPaginationParams(pageNumber, pageSize)

		todos, todoErr := controller.model.PaginateTodo(
			user.Id,
			pagination.PageSize,
			(pagination.PageNumber-1)*pagination.PageSize)

		if todoErr != nil {
			log.Println(todoErr)
			context.JSON(FailureMessageResponse(todoErr.Error()))
			return
		}

		count, countErr := controller.model.CountPaginationTodo(user.Id)

		if countErr != nil {
			log.Println(countErr)
			count = 0
		}

		// log.Println(todos)
		context.JSON(SuccessResponse(TODOS_READ_SUCCESSFULLY, T{
			"rows": todos,
			"pagination": GetPaginationParams(
				count, pagination.PageNumber, pagination.PageSize),
		}))
	}
}

func (controller *TodoController) UpdateTodoTask() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		if !isUser {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		// Get todo id from params
		var todoId, idErr = ConvertStringIdToInt(context.Param("id"))

		if idErr != nil {
			log.Println(idErr)
			context.JSON(FailureMessageResponse(INVALID_ID))
			return
		}

		// Get todo task from request body
		var payload CreateTodoRequestBody

		if err := context.BindJSON(&payload); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		validatedString, isValid := ValidateString(payload.Task)

		if !isValid {
			context.JSON(FailureMessageResponse(INVALID_TODO))
			return
		}

		// Update todo
		if err := controller.model.UpdateTodoTask(user.Id, todoId, validatedString); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(SuccessMessageResponse(TODO_UPDATED_SUCCESSFULLY))
		return
	}
}

func (controller *TodoController) UpdateTodoCompleted() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		if !isUser {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		// Get todo id from params
		var todoId, idErr = ConvertStringIdToInt(context.Param("id"))

		if idErr != nil {
			log.Println(idErr)
			context.JSON(FailureMessageResponse(INVALID_ID))
			return
		}

		// Get todo state from request body
		var payload UpdateTodoStateRequestBody

		if err := context.BindJSON(&payload); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		if fmt.Sprintf("%T", payload.Completed) != "bool" {
			context.JSON(FailureMessageResponse(INVALID_TODO_STATE))
			return
		}

		// Update todo
		err := controller.model.UpdateTodoCompletedState(
			user.Id, todoId, payload.Completed)

		if err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(SuccessMessageResponse(TODO_UPDATED_SUCCESSFULLY))
		return
	}
}

func (controller *TodoController) DeleteTodo() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		if !isUser {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		// Get todo id from params
		var todoId, idErr = ConvertStringIdToInt(context.Param("id"))

		if idErr != nil {
			log.Println(idErr)
			context.JSON(FailureMessageResponse(INVALID_ID))
			return
		}

		if err := controller.model.DeleteTodo(user.Id, todoId); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(SuccessMessageResponse(TODO_DELETED_SUCCESSFULLY))
	}
}
