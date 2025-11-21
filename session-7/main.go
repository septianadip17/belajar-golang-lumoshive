package main

import (
	"session-7/dto"
	"session-7/handler"
	"session-7/service"
)

func main() {
	requestRegister := dto.RegisterUserRequest{
		Name:  "iay",
		Email: "iay@gmail.com",
		Age:   20,
	}

	// init
	service := service.NewUserService()
	handler := handler.NewUserHandler(service)

	handler.Register(requestRegister)
}
