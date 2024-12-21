package service

import (
	"go-template/internal/identity/core/model"
)

type UpdateUserInput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *Service) UpdateUser(input UpdateUserInput) (*model.User, error) {
	user, err := s.userRepository.Update(input.ID, input.Name, input.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
