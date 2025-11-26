package service

import (
	"errors"
	"session-7/dto"
	"session-7/model"
	"session-7/utils"
	"strings"
	"time"
)

type UserService struct{}

// constractor
func NewUserService() UserService {
	return UserService{}
}

// method register user
func (userservice *UserService) RegisterUser(req dto.RegisterUserRequest) (model.User, error) {

	// Validasi sederhana
	if strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.Email) == "" {
		return model.User{}, errors.New("name, and email are required")
	}

	// Baca data existing
	users, err := utils.ReadUsersFromFile()
	if err != nil {
		return model.User{}, err
	}

	// Cek email duplikat
	for _, u := range users {
		if u.Email == req.Email {
			return model.User{}, errors.New("email already registered")
		}
	}

	// Generate ID sederhana (max ID + 1)
	newID := 1
	for _, u := range users {
		if int(u.Id) >= newID {
			newID = u.Id + 1
		}
	}

	// create password
	password := "123"

	newUser := model.User{
		Base: model.Base{
			Id:        newID,
			CreatedAt: time.Now(),
		},
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		Password: password,
	}

	users = append(users, newUser)

	// save data
	if err := utils.WriteUsersToFile(users); err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

// method list user
func (userservice *UserService) ListUser() (*[]dto.ListUserResponse, error) {

	// read data existing
	users, err := utils.ReadUsersFromFile()
	if err != nil {
		return nil, err
	}

	var listUsers []dto.ListUserResponse

	for _, u := range users {
		user := dto.ListUserResponse{
			Name:  u.Name,
			Email: u.Email,
		}
		listUsers = append(listUsers, user)
	}

	return &listUsers, nil
}
