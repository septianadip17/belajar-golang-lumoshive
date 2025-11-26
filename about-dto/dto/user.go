package dto

type RegisterUserRequest struct {
	Name  string
	Email string
	Age   int
}

type ListUserResponse struct {
	Name  string
	Email string
}
