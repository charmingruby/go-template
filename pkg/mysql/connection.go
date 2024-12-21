package mysql

import (
	"database/sql"
	"fmt"
	"go-template/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func New(envs *config.Environment) (*sql.DB, error) {
	USER := envs.DATABASE_USER
	PASSWORD := envs.DATABASE_PASSWORD
	HOST := envs.DATABASE_HOST
	PORT := envs.DATABASE_PORT
	NAME := envs.DATABASE_NAME

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASSWORD, HOST, PORT, NAME))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return db, nil
}
