package utils

import (
	"encoding/json"
	"errors"
	"os"
	"session-7/model"
)

const UsersFilePath = "data/users.json"

// cek directory and file
func EnsureUsersFile() error {
	_, err := os.Stat(UsersFilePath)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll("data", 0755); err != nil {
			return err
		}
		return os.WriteFile(UsersFilePath, []byte("[]"), 0644)
	}
	return nil
}

// read file
func ReadUsersFromFile() ([]model.User, error) {
	if err := EnsureUsersFile(); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(UsersFilePath)
	if err != nil {
		return nil, err
	}

	var users []model.User
	if err := json.Unmarshal(bytes, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// write file
func WriteUsersToFile(users []model.User) error {
	bytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(UsersFilePath, bytes, 0644)
}
