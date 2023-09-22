package web

import (
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
	"github.com/otumian-empire/go-to-do-list-restful-api/repository"
)

type TodoController struct {
	model repository.Repository
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

func (controller *TodoController) ReadTodo() gin.HandlerFunc {
	return func(context *gin.Context) {
		value, isValue := context.MustGet("user").(model.User)

		log.Println(value)

		if !isValue {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		var id = context.Param("id")

		if len(id) < 1 {
			context.JSON(FailureMessageResponse(INVALID_ID))
			return
		}

		intId, intIdErr := strconv.Atoi(id)

		if intIdErr != nil {
			context.JSON(FailureMessageResponse(INVALID_ID))
			return
		}

		todo, todoErr := controller.model.ReadTodoById(value.Id, intId)

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
		value, isValue := context.MustGet("user").(model.User)

		if !isValue {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		pageNumber, pageNumberErr := strconv.Atoi(
			context.DefaultQuery("pageNumber", "1"))

		if pageNumberErr != nil {
			pageNumber = DEFAULT_PAGE_NUMBER
		}

		pageSize, pageSizeErr := strconv.Atoi(
			context.DefaultQuery("pageSize", "20"))

		if pageSizeErr != nil {
			pageSize = DEFAULT_PAGE_SIZE
		}

		var pagination = CleanPaginationParams(pageNumber, pageSize)

		todos, todoErr := controller.model.PaginateTodo(
			value.Id,
			pagination.PageSize,
			(pagination.PageNumber-1)*pagination.PageSize)

		if todoErr != nil {
			log.Println(todoErr)
			context.JSON(FailureMessageResponse(todoErr.Error()))
			return
		}

		count, countErr := controller.model.CountPaginationTodo(value.Id)

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
		value, isValue := context.MustGet("user").(model.User)

		if !isValue {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		log.Println(value)

	}
}

// func (controller *TodoController) UpdateTodoCompleted() gin.HandlerFunc

// func (controller *TodoController) DeleteTodo() gin.HandlerFunc
