package service

import (
	"go-template/internal/identity/core/model"
)

func (s *Service) FindAllUsers() ([]model.User, error) {
	users, err := s.userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
