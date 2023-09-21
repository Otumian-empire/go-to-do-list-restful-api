package web

// use for both sign up and login
type AuthRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// use for update username
type UpdateUserUsernameRequestBody struct {
	Username string `json:"username"`
}
