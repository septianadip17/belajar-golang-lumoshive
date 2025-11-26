package main

import (
	"session-7/handler"
	"session-7/service"
)

func main() {

	// // inputan form user
	// requestRegister := dto.RegisterUserRequest{
	// 	Name:  "academy",
	// 	Email: "sidiklumosacademy@gmail.com",
	// 	Age:   30,
	// }

	// init
	service := service.NewUserService()
	handler := handler.NewUserHandler(service) // dependency inject

	// method register
	// handler.Register(requestRegister)
	handler.ListUser()
}
