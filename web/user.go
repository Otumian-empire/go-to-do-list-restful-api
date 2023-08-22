package web

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
)

type UserController struct {
	model model.UserModel
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

		// TODO: install bcrypt and create a hash of the password

		// create user
		if err := controller.model.CreateUser(payload.Username, payload.Password); err != nil {
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

		if err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		// TODO: compare the user password passed to the hash
		row.Password = ""

		// return success message on user creation
		log.Println(LOGIN_SUCCESSFUL)
		context.JSON(SuccessResponse(LOGIN_SUCCESSFUL, row))
	}
}

func (controller *UserController) UpdateUserUsername() gin.HandlerFunc

func (controller *UserController) UpdateUserPassword() gin.HandlerFunc

func (controller *UserController) ReadUser() gin.HandlerFunc

func (controller *UserController) DeleteUser() gin.HandlerFunc

func (controller *UserController) Logout() gin.HandlerFunc
