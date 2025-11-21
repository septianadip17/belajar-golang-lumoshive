package handler

import {
	"session-7/service"
}


type UserHandler struct{
	Userservice service.UserService
}

// constractor
func NewUserHandler() UserHandler {
	return UserHandler{}
}

func (*userHandler) RegisterUser() (model.user, error) {
	return model.user{}, nil
}

