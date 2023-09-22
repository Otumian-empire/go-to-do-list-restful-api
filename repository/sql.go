package repository

// User queries
const (
	CREATE_USER_QUERY              = "INSERT INTO `user`(`username`, `password`) VALUES (?, ?)"
	READ_USER_BY_ID_QUERY          = "SELECT `id`, `username`, `password`, `created_at`, `updated_at` FROM `user` WHERE `id`=?"
	GET_USERNAME_BY_USERNAME_QUERY = "SELECT `id`, `username`, `password`, `created_at`, `updated_at` FROM `user` WHERE `username`=?"
	UPDATE_PASSWORD_QUERY          = "UPDATE `user` SET `password`=? WHERE `id`=?"
	UPDATE_USERNAME_QUERY          = "UPDATE `user` SET `username`=? WHERE `id`=?"
	DELETE_USER_QUERY              = "DELETE FROM `user` WHERE `id`=?"
)

// Todo queries
const (
	CREATE_TODO_QUERY                 = "INSERT INTO `todo`(`task`, `completed`, `user`) VALUES (?,?,?)"
	READ_TODO_BY_ID_QUERY             = "SELECT `id`, `task`, `completed`, `created_at`, `updated_at`, `user` FROM `todo` WHERE `user`=? AND `id`=?"
	PAGINATE_TODO_QUERY               = "SELECT `id`, `task`, `completed`, `created_at`, `updated_at`, `user` FROM `todo` WHERE `user`=? LIMIT ? OFFSET ?"
	PAGINATION_TODO_COUNT_QUERY       = "SELECT COUNT(*) as `count` FROM `todo` WHERE `user`=?"
	UPDATE_TODO_TASK_QUERY            = "UPDATE `todo` SET `task`=? WHERE `user`=? AND `id`=?"
	UPDATE_TODO_COMPLETED_STATE_QUERY = "UPDATE `todo` SET `completed`=? WHERE `user`=? AND `id`=?"
	DELETE_TODO_QUERY                 = "DELETE FROM `todo` WHERE `user`=? AND `id`=?"
	DELETE_TODOS_QUERY                = "DELETE FROM `todo` WHERE `user`=?"
)
