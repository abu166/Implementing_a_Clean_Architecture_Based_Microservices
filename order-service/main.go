package main

import (
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/config"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/database"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()
    database.InitDB()

    r := gin.Default()

    // Routes
    r.POST("/orders", handlers.CreateOrder)
    r.GET("/orders/:id", handlers.GetOrder)
    r.PATCH("/orders/:id", handlers.UpdateOrderStatus)
    r.GET("/orders", handlers.ListOrders)

    r.Run(":8082")
}