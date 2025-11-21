package model

type User struct {
	Base
	Name     string
	Age      int
	Email    string
	Password string
}
