package web

type AuthRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserUsernameRequestBody struct {
	Username string `json:"username"`
}

type UpdateUserPasswordRequestBody struct {
	Password string `json:"password"`
}

type CreateTodoRequestBody struct {
	Task string `json:"task"`
}
