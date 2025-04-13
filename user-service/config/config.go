package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DbHost     string
    DbPort     string
    DbUser     string
    DbPassword string
    DbName     string
}

var AppConfig *Config

func LoadConfig() {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    AppConfig = &Config{
        DbHost:     os.Getenv("DB_HOST"),
        DbPort:     os.Getenv("DB_PORT"),
        DbUser:     os.Getenv("DB_USER"),
        DbPassword: os.Getenv("DB_PASSWORD"),
        DbName:     os.Getenv("DB_NAME"),
    }
}