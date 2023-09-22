package web

import (
	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/repository"
)

func NewHandler(_repository repository.Repository, router *gin.Engine) *gin.Engine {

	// Call the controllers and pass the repository
	userController := UserController{model: _repository}
	todoController := TodoController{model: _repository}

	// Endpoints specific to create, read and delete of the url
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.SignUp())
		userRoutes.POST("/auth", userController.Login())

		userRoutes.Use(ApiKeyAuth(_repository))

		userRoutes.PUT("/username", userController.UpdateUserUsername())
		userRoutes.PUT("/password", userController.UpdateUserPassword())
		userRoutes.GET("/", userController.ReadUser())
		userRoutes.DELETE("/", userController.DeleteUser())
		// userRoutes.GET("/logout", userController.Logout())
	}

	// Endpoints specific for todos
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.Use(ApiKeyAuth(_repository))

		todoRoutes.POST("/", todoController.CreateTodo())
		todoRoutes.GET("/", todoController.ReadTodos())
		todoRoutes.GET("/:id", todoController.ReadTodo())
		todoRoutes.PUT("/:id", todoController.UpdateTodoTask())
		todoRoutes.PUT("/:id/state", todoController.UpdateTodoCompleted())
		todoRoutes.DELETE("/:id", todoController.DeleteTodo())
	}

	return router
}
