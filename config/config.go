package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DBConfig
	App AppConfig
}

type DBConfig struct {
	Name     string
	Password string
	User     string
	Host     string
	Port     string
	DSN      string
}

type AppConfig struct {
	Port string
}

func New() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	dbName := os.Getenv("POSTGRES_DB")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbUser := os.Getenv("POSTGRES_USER")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	appPort := os.Getenv("APP_PORT")

	dsnFormat := "postgresql://%s:%s@%s:%s/%s?sslmode=disable"
	dsn := fmt.Sprintf(dsnFormat, dbUser, dbPassword, dbHost, dbPort, dbName)

	dbConfig := DBConfig{
		Name:     dbName,
		Password: dbPassword,
		User:     dbUser,
		Host:     dbHost,
		Port:     dbPort,
		DSN:      dsn,
	}

	appConfig := AppConfig{
		Port: appPort,
	}

	cnf := Config{
		DB:  dbConfig,
		App: appConfig,
	}

	return cnf, nil
}
