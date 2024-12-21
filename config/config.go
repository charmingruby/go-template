package config

import (
	"log"
	"os"
)

type Environment struct {
	DATABASE_HOST     string
	DATABASE_NAME     string
	DATABASE_PASSWORD string
	DATABASE_PORT     string
	DATABASE_USER     string
	PORT              string
}

func New() *Environment {

	return &Environment{
		DATABASE_HOST:     getEnvOrDie("DATABASE_HOST"),
		DATABASE_NAME:     getEnvOrDie("DATABASE_NAME"),
		DATABASE_PASSWORD: getEnvOrDie("DATABASE_PASSWORD"),
		DATABASE_PORT:     getEnvOrDie("DATABASE_PORT"),
		DATABASE_USER:     getEnvOrDie("DATABASE_USER"),
		PORT:              getEnvOrDie("PORT"),
	}
}

func getEnvOrDie(env string) (val string) {
	val = os.Getenv(env)

	if val == "" {
		log.Fatalf("missing environment variable %s", env)
	}

	return val
}
