package requests

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}