package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV        string
	HTTPServer *HTTPServer
	Database   *Database
}

type HTTPServer struct {
	Address     string
	IDLETimeout time.Duration
}

type Database struct {
	URL      string
	User     string
	Password string
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Panic("error loading .env file")
	}

	env := os.Getenv("ENV")
	if env == "" {
		log.Panic("error: ENV is empty")
	}

	databaseURL := os.Getenv("POSTGRES_URL")
	if databaseURL == "" {
		log.Panic("error: POSTGRES_URL is empty")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		log.Panic("error: POSTGRES_USER is empty")
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		log.Panic("error: POSTGRES_PASSWORD is empty")
	}

	serverAddress := os.Getenv("HTTP_SERVER_ADDRESS")
	if password == "" {
		log.Panic("error: HTTP_SERVER_ADDRESS not set")
	}

	serverIDLE := os.Getenv("HTTP_SERVER_IDLE")
	if password == "" {
		log.Panic("error: HTTP_SERVER_IDLE not set")
	}

	idle, err := strconv.Atoi(serverIDLE)
	if err != nil {
		log.Panicf("%w", err)
	}

	return &Config{
		ENV: env,
		HTTPServer: &HTTPServer{
			Address:     serverAddress,
			IDLETimeout: time.Duration(idle),
		},
		Database: &Database{
			URL:      databaseURL,
			User:     user,
			Password: password,
		},
	}
}
