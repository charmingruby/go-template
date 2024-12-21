package model

import "github.com/oklog/ulid/v2"

func NewUser(name, email string) *User {
	return &User{
		ID:    ulid.Make().String(),
		Email: email,
		Name:  name,
	}
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
