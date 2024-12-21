package identity

import (
	"database/sql"
	"go-template/internal/identity/core/repository/mysql_repository"
	"go-template/internal/identity/core/service"
	"go-template/internal/identity/transport/rest/endpoint"

	"github.com/gofiber/fiber/v2"
)

func NewUserRepository(db *sql.DB) (*mysql_repository.UserRepository, error) {
	return mysql_repository.NewUserRepository(db)
}

func NewService(userRepo *mysql_repository.UserRepository) *service.Service {
	return service.New(userRepo)
}

func NewRestHandler(router *fiber.App, service service.Service) {
	endpoint.New(service).Setup(router)
}
