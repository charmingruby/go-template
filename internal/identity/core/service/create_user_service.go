package service

import (
	"go-template/internal/identity/core/model"
)

type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *Service) CreateUser(input CreateUserInput) error {
	user := model.NewUser(input.Name, input.Email)

	if err := s.userRepository.Create(*user); err != nil {
		return err
	}

	return nil
}
