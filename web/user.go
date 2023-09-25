package web

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
	"github.com/otumian-empire/go-to-do-list-restful-api/repository"
)

type UserController struct {
	model repository.Repository
}

func (controller *UserController) SignUp() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get the data from the request body
		var payload AuthRequestBody

		if err := context.BindJSON(&payload); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		validatedUsername, isValidUsername := ValidateString(payload.Username)
		if !isValidUsername {
			context.JSON(FailureMessageResponse(INVALID_USERNAME))
			return
		}

		validatedPassword, isValidPassword := ValidateString(payload.Password)
		if !isValidPassword {
			context.JSON(FailureMessageResponse(INVALID_PASSWORD))
			return
		}

		passwordHash, err := HashPassword(validatedPassword)

		if err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		// Create user
		if err := controller.model.User.CreateUser(validatedUsername, passwordHash); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		// Return success message on user creation
		context.JSON(SuccessMessageResponse(SIGN_UP_SUCCESSFUL))
	}
}

func (controller *UserController) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get the data from the request body
		var payload AuthRequestBody

		if err := context.BindJSON(&payload); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		validatedUsername, isValidUsername := ValidateString(payload.Username)
		if !isValidUsername {
			context.JSON(FailureMessageResponse(INVALID_USERNAME))
			return
		}

		validatedPassword, isValidPassword := ValidateString(payload.Password)
		if !isValidPassword {
			context.JSON(FailureMessageResponse(INVALID_PASSWORD))
			return
		}

		// Read row with the same username
		user, err := controller.model.ReadUserByUsername(validatedUsername)

		if err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		if !CheckPasswordHash(validatedPassword, user.Password) {
			log.Println("Validation check failed")
			context.JSON(FailureMessageResponse(INVALID_CREDENTIAL))
			return
		}

		token, tokenErr := JWTAuthService().TokenGenerate(user.Id)

		if tokenErr != nil {
			log.Println(tokenErr)
			context.JSON(FailureMessageResponse(tokenErr.Error()))
			return
		}

		log.Println("token:", token)

		// Return success message on user creation
		context.JSON(SuccessResponse(LOGIN_SUCCESSFUL, T{
			"user": T{
				"id":        user.Id,
				"username":  user.Username,
				"createdAt": user.CreatedAt,
				"updatedAt": user.UpdatedAt,
			},
			"token": token,
		}))
	}
}

func (controller *UserController) UpdateUserUsername() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		if !isUser {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		// Get the username from the request body
		var payload UpdateUserUsernameRequestBody

		// Check if the user is valid
		if err := context.BindJSON(&payload); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		validatedUsername, isValidUsername := ValidateString(payload.Username)

		if !isValidUsername {
			log.Println(INVALID_USERNAME)
			context.JSON(FailureMessageResponse(INVALID_USERNAME))
			return
		}

		// Query the database for a row that matches the new username
		// If the error is not <nil> the there is a row with the same username
		if _, err := controller.model.ReadUserByUsername(validatedUsername); err == nil {
			context.JSON(FailureMessageResponse(USERNAME_TAKEN))
			return
		}

		// Update the username
		if err := controller.model.UpdateUserUsername(user.Id, validatedUsername); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(COULD_NOT_UPDATE_USERNAME))
			return
		}

		// Return a success response
		context.JSON(SuccessMessageResponse(USERNAME_UPDATED_SUCCESSFULLY))
	}
}

func (controller *UserController) UpdateUserPassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		if !isUser {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		// Get the password from the request body
		var payload UpdateUserPasswordRequestBody

		// Check if the user is valid
		if err := context.BindJSON(&payload); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(err.Error()))
			return
		}

		validatedPassword, isValidPassword := ValidateString(payload.Password)

		if !isValidPassword {
			log.Println(INVALID_PASSWORD)
			context.JSON(FailureMessageResponse(INVALID_PASSWORD))
			return
		}

		passwordHash, passwordHashingError := HashPassword(validatedPassword)

		if passwordHashingError != nil {
			log.Println(passwordHashingError)
			context.JSON(FailureMessageResponse(passwordHashingError.Error()))
			return
		}

		// Update the password
		if err := controller.model.UpdateUserPassword(user.Id, passwordHash); err != nil {
			log.Println(err.Error())
			context.JSON(FailureMessageResponse(COULD_NOT_UPDATE_PASSWORD))
			return
		}

		// Return a success response
		context.JSON(SuccessMessageResponse(PASSWORD_UPDATED_SUCCESSFULLY))
	}
}

func (controller *UserController) ReadUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		log.Println("user", user)
		log.Println("isUser", isUser)

		if !isUser {
			context.JSON(AuthenticationErrorResponse(INVALID_AUTHENTICATION))
			return
		}

		context.JSON(SuccessResponse(USER_DETAIL_READ_SUCCESSFULLY, T{
			"id":        user.Id,
			"username":  user.Username,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		}))
	}
}

func (controller *UserController) DeleteUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isUser := context.MustGet("user").(model.User)

		if !isUser {
			context.JSON(FailureMessageResponse(INVALID_AUTHENTICATION))
			return
		}

		if err := controller.model.DeleteTodos(user.Id); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(COULD_NOT_DELETE_USER))
			return
		}

		if err := controller.model.DeleteUser(user.Id); err != nil {
			log.Println(err)
			context.JSON(FailureMessageResponse(COULD_NOT_DELETE_USER))
			return
		}

		context.JSON(SuccessMessageResponse(USER_DELETED_SUCCESSFULLY))
	}
}

// func (controller *UserController) Logout() gin.HandlerFunc
