package web

const (
	USERNAME_UPDATED_SUCCESSFULLY = "username updated successfully"
	PASSWORD_UPDATED_SUCCESSFULLY = "password updated successfully"
	USER_DETAIL_READ_SUCCESSFULLY = "user detail read successfully"
	USER_DELETED_SUCCESSFULLY     = "user deleted successfully"
	TODO_CREATED_SUCCESSFULLY     = "todo created successfully"
	TODO_READ_SUCCESSFULLY        = "todo read successfully"
	TODOS_READ_SUCCESSFULLY       = "todos read successfully"
	TODO_UPDATED_SUCCESSFULLY     = "todo updated successfully"

	INVALID_USERNAME   = "invalid username"
	INVALID_PASSWORD   = "invalid password"
	INVALID_TODO       = "invalid todo"
	INVALID_TODO_STATE = "invalid todo state, pass true or false"
	INVALID_ID         = "invalid id"

	USERNAME_TAKEN            = "username is taken, try another one"
	COULD_NOT_UPDATE_USERNAME = "could not update username"
	COULD_NOT_UPDATE_PASSWORD = "could not update password"
	COULD_NOT_DELETE_USER     = "could not delete user"
)

const (
	SIGN_UP_SUCCESSFUL = "user created successfully, please login"
	LOGIN_SUCCESSFUL   = "login successfully"

	INVALID_AUTHENTICATION = "invalid authentication"
	INVALID_CREDENTIAL     = "incorrect username or password"
)

const (
	DATABASE_CONNECTED               = "database connected"
	SERVER_RECOVER_FROM_ERROR        = "recover from an error"
	SERVER_LOADING_CREDENTIALS_ERROR = "error loading .env file"
	SERVER_RUNNING_ON_PORT           = "server running on port"
)
