package web

const (
	SIGN_UP_SUCCESSFUL            = "user created successfully, please login"
	LOGIN_SUCCESSFUL              = "login successfully"
	USERNAME_UPDATED_SUCCESSFULLY = "username updated successfully"
	PASSWORD_UPDATED_SUCCESSFULLY = "password updated successfully"
	USER_DETAIL_READ_SUCCESSFULLY = "user detail read successfully"
	USER_DELETED_SUCCESSFULLY     = "user deleted successfully"
	TODO_CREATED_SUCCESSFULLY     = "todo created successfully"

	INVALID_CREDENTIAL        = "incorrect username or password"
	INVALID_USERNAME          = "invalid username"
	INVALID_PASSWORD          = "invalid password"
	INVALID_TODO              = "invalid todo"
	USERNAME_TAKEN            = "username is taken, try another one"
	COULD_NOT_UPDATE_USERNAME = "could not update username"
	COULD_NOT_UPDATE_PASSWORD = "could not update password"
	COULD_NOT_DELETE_USER     = "could not delete user"

	DATABASE_CONNECTED = "database connected"
)

const (
	INVALID_AUTHENTICATION = "invalid authentication"
)

const (
	SERVER_RECOVER_FROM_ERROR        = "recover from an error"
	SERVER_LOADING_CREDENTIALS_ERROR = "error loading .env file"
	SERVER_RUNNING_ON_PORT           = "server running on port"
)
