package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/config"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/models"
)

var DB *gorm.DB // Import gorm package

func InitDB() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        config.AppConfig.DbHost, config.AppConfig.DbUser, config.AppConfig.DbPassword, config.AppConfig.DbName, config.AppConfig.DbPort)
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Use the gorm package to open a connection
    if err != nil {
        panic("Failed to connect to the database")
    }

    // Auto migrate the Order and OrderItem models
    DB.AutoMigrate(&models.Order{}, &models.OrderItem{})
}