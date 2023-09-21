package web

import (
	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/repository"
)

func NewHandler(_repository repository.Repository, router *gin.Engine) *gin.Engine {

	// call the controllers and pass the repository
	userController := UserController{model: _repository}
	// add other controllers here

	// endpoints specific to create, read and delete of the url
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.SignUp())
		userRoutes.POST("/auth", userController.Login())

		userRoutes.Use(ApiKeyAuth(_repository))

		userRoutes.PUT("/username", userController.UpdateUserUsername())
		// userRoutes.PUT("/password", userController.UpdateUserPassword())
		// userRoutes.GET("/", userController.ReadUser())
		// userRoutes.DELETE("/", userController.DeleteUser())
		// userRoutes.GET("/logout", userController.Logout())
	}

	// endpoint to redirect to the actual url
	// router.GET("/:hash", userController.GerOriginalUrl())

	return router
}
