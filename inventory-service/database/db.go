package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/config"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/models"
)

var DB *gorm.DB

func InitDB() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        config.AppConfig.DbHost, config.AppConfig.DbUser, config.AppConfig.DbPassword, config.AppConfig.DbName, config.AppConfig.DbPort)
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }

    // Auto migrate the Product model
    DB.AutoMigrate(&models.Product{})
}