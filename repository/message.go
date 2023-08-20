package repository

// User defined global
const (
	GLOBAL_ERROR        = "an error occurred please contact support"
	NO_ROW_AFFECT_ERROR = "no rows affected"
	NO_ROW_FOUND        = "no row found"
)

// User defined error messages for database connection
const (
	DATABASE_OPENING_ERROR    = ""
	DATABASE_CONNECTING_ERROR = ""
)

// User defined error messages for the user model
const (
	CREATE_USER_ERROR = "could not create user"
	DELETE_USER_ERROR = "could not delete user"
	UPDATE_USER_ERROR = "could not update user"
)

// User defined error messages for the todo model
const (
	CREATE_TODO_ERROR   = "could not create todo"
	DELETE_TODO_ERROR   = "could not delete todo"
	UPDATE_TODO_ERROR   = "could not update todo"
	PAGINATE_TODO_ERROR = "could not read todos"
)
