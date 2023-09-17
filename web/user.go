package web

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/repository"
)

type UserController struct {
	model repository.Repository
}

func (controller *UserController) SignUp() gin.HandlerFunc {
	return func(context *gin.Context) {
		// get the data from the request body
		var payload AuthRequestBody

		if err := context.BindJSON(&payload); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		passwordHash, err := HashPassword(payload.Password)

		if err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		payload.Password = passwordHash

		// create user
		if err := controller.model.User.CreateUser(payload.Username, payload.Password); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		// return success message on user creation
		log.Println(SIGN_UP_SUCCESSFUL)
		context.JSON(SuccessMessageResponse(SIGN_UP_SUCCESSFUL))
	}
}

func (controller *UserController) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		// get the data from the request body
		var payload AuthRequestBody

		if err := context.BindJSON(&payload); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		// read row with the same username
		row, err := controller.model.ReadUserByUsername(payload.Username)

		log.Println(row)
		if err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		if !CheckPasswordHash(payload.Password, row.Password) {
			log.Println("Validation check failed")
			context.JSON(FailureMessageResponse(INVALID_CREDENTIAL))
			return
		}

		row.Password = ""

		// return success message on user creation
		log.Println(LOGIN_SUCCESSFUL)
		context.JSON(SuccessResponse(LOGIN_SUCCESSFUL, row))
	}
}

// func (controller *UserController) UpdateUserUsername() gin.HandlerFunc

// func (controller *UserController) UpdateUserPassword() gin.HandlerFunc

// func (controller *UserController) ReadUser() gin.HandlerFunc

// func (controller *UserController) DeleteUser() gin.HandlerFunc

// func (controller *UserController) Logout() gin.HandlerFunc
