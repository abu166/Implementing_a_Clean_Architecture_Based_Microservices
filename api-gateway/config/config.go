package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    InventoryServiceURL string
    OrderServiceURL     string
}

var AppConfig *Config

func LoadConfig() {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    AppConfig = &Config{
        InventoryServiceURL: os.Getenv("INVENTORY_SERVICE_URL"),
        OrderServiceURL:     os.Getenv("ORDER_SERVICE_URL"),
    }
}