package web

const (
	SIGN_UP_SUCCESSFUL            = "user created successfully, please login"
	LOGIN_SUCCESSFUL              = "login successfully"
	USERNAME_UPDATED_SUCCESSFULLY = "username updated successfully"

	INVALID_CREDENTIAL        = "Incorrect username or password"
	INVALID_USERNAME          = "Invalid username"
	USERNAME_TAKEN            = "Username is taken, try another one"
	COULD_NOT_UPDATE_USERNAME = "could not update username"

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
