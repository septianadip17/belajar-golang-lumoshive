package dto

import "session-7/model"

type RegisterUserRequest struct {
	Name     string
	Age      int
	Email    string
}

type RegisterUserResponse