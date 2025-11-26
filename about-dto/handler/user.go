package handler

import (
	"encoding/json"
	"fmt"
	"session-7/dto"
	"session-7/service"
)

type UserHandler struct {
	Userservice service.UserService
}

// constractor
func NewUserHandler(userservice service.UserService) UserHandler {
	return UserHandler{
		Userservice: userservice,
	}
}

// method register user
func (userhandler *UserHandler) Register(req dto.RegisterUserRequest) {

	user, err := userhandler.Userservice.RegisterUser(&req)

	if err != nil {
		fmt.Println("error unregister")
	}

	fmt.Printf("Create user success , User : %+v\n", user)
}

// method list user
func (userhandler *UserHandler) ListUser() {
	users, err := userhandler.Userservice.ListUser()
	if err != nil {
		fmt.Println("error data")
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("error marshalling json")
	}

	fmt.Println(string(data))
}
