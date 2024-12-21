package service

import (
	"go-template/internal/identity/core/repository/mysql_repository"
)

type Service struct {
	userRepository *mysql_repository.UserRepository
}

func New(userRepository *mysql_repository.UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}
