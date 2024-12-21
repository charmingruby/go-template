package main

import (
	"database/sql"
	"fmt"
	"go-template/config"
	"go-template/internal/identity"
	"go-template/internal/shared/transport/rest"
	"go-template/pkg/mysql"
	"go-template/pkg/shutdown"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("error loading .env file")
	}

	defer shutdown.Gracefully()

	config := config.New()

	db, err := mysql.New(config)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}

	router := rest.New()

	if err := initDependencies(db, router); err != nil {
		log.Fatalf("error initializing dependencies: %s", err)
	}

	if err := router.Listen(fmt.Sprintf(":%s", config.PORT)); err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}

func initDependencies(db *sql.DB, router *fiber.App) error {
	userRepo, err := identity.NewUserRepository(db)
	if err != nil {
		return err
	}

	identitySvc := identity.NewService(userRepo)

	identity.NewRestHandler(router, *identitySvc)

	return nil
}
