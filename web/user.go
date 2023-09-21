package web

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
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

		// generate authorization token
		// token, tokenErr := /* GenerateJwt */ CreateToken(
		// 	CustomPayload{
		// 		Id:       row.Id,
		// 		Username: row.Username,
		// 	},
		// )

		// if tokenErr != nil {
		// 	log.Println("Token generation error")
		// 	context.JSON(FailureMessageResponse(tokenErr.Error()))
		// 	return
		// }

		// return success message on user creation
		log.Println(LOGIN_SUCCESSFUL)
		context.JSON(SuccessResponse(LOGIN_SUCCESSFUL, T{
			"user": T{
				"id":        row.Id,
				"username":  row.Username,
				"createdAt": row.CreatedAt,
				"updatedAt": row.UpdatedAt,
			},
			// "token": token,
		}))

		return
	}
}

func (controller *UserController) UpdateUserUsername() gin.HandlerFunc {
	return func(context *gin.Context) {
		value, isValue := context.MustGet("user").(model.User)

		log.Println(value)

		if !isValue {
			context.JSON(SuccessMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		// get the username from the request body
		var payload UpdateUserUsernameRequestBody

		// check if the user is valid
		if err := context.BindJSON(&payload); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		payload.Username = strings.Trim(payload.Username, " ")

		if len(payload.Username) < 1 {
			log.Println(INVALID_USERNAME)
			context.JSON(FailureMessageResponse(INVALID_USERNAME))
			return
		}

		// query the database for a row that matches the new username
		if _, err := controller.model.ReadUserByUsername(payload.Username); err == nil {
			// check that the user's username is the same as that of the auth user
			// else user already exist
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(USERNAME_TAKEN))
			return
		}

		// update the username
		if err := controller.model.UpdateUserUsername(value.Id, payload.Username); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(COULD_NOT_UPDATE_USERNAME))
			return
		}

		// return a success response
		context.JSON(SuccessMessageResponse(USERNAME_UPDATED_SUCCESSFULLY))
	}
}

func (controller *UserController) UpdateUserPassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		value, isValue := context.MustGet("user").(model.User)

		log.Println(value)

		if !isValue {
			context.JSON(SuccessMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		// get the password from the request body
		var payload UpdateUserPasswordRequestBody

		// check if the user is valid
		if err := context.BindJSON(&payload); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		payload.Password = strings.Trim(payload.Password, " ")

		if len(payload.Password) < 1 {
			log.Println(INVALID_PASSWORD)
			context.JSON(FailureMessageResponse(INVALID_PASSWORD))
			return
		}

		passwordHash, passwordHashingError := HashPassword(payload.Password)

		if passwordHashingError != nil {
			log.Println(passwordHashingError)
			context.JSON(SuccessMessageResponse(passwordHashingError.Error()))
			return
		}

		// update the password
		if err := controller.model.UpdateUserPassword(value.Id, passwordHash); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(COULD_NOT_UPDATE_PASSWORD))
			return
		}

		// return a success response
		context.JSON(SuccessMessageResponse(PASSWORD_UPDATED_SUCCESSFULLY))
	}
}

// func (controller *UserController) ReadUser() gin.HandlerFunc

// func (controller *UserController) DeleteUser() gin.HandlerFunc

// func (controller *UserController) Logout() gin.HandlerFunc
