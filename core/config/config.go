package config

import (
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	KINOPOISK_KEY, DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, MIGRATION_PREFIX, PB_TARGET string
	JWT_TOKEN                                                                                   []byte
}

var Cfg Config

func Initialize(envPath string) {

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("couldnt load .env", "err", err)
	}

	Cfg = Config{
		KINOPOISK_KEY:    os.Getenv("KINOPOISK_KEY"),
		DB_USER:          os.Getenv("DB_USER"),
		DB_PASSWORD:      os.Getenv("DB_PASSWORD"),
		DB_HOST:          os.Getenv("DB_HOST"),
		DB_PORT:          os.Getenv("DB_PORT"),
		DB_NAME:          os.Getenv("DB_NAME"),
		MIGRATION_PREFIX: os.Getenv("MIGRATION_PREFIX"),
		PB_TARGET:        os.Getenv("PB_TARGET"),
		JWT_TOKEN:        []byte(os.Getenv("JWT_TOKEN")),
	}

}
