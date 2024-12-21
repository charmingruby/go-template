package service

import (
	"go-template/internal/identity/core/model"
)

type FindOneUserInput struct {
	ID string `json:"id"`
}

func (s *Service) FindOneUser(input FindOneUserInput) (*model.User, error) {
	user, err := s.userRepository.FindOne(input.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
