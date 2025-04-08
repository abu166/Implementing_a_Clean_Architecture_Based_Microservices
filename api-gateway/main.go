package main

import (
    "github.com/gin-gonic/gin"
    "Implementing_a_Clean_Architecture_Based_Microservices/api-gateway/handlers"
    "Implementing_a_Clean_Architecture_Based_Microservices/api-gateway/middleware"
    "Implementing_a_Clean_Architecture_Based_Microservices/api-gateway/config"
)

func main() {
    config.LoadConfig()

    r := gin.Default()

    // Middleware
    r.Use(middleware.LoggingMiddleware())
    r.Use(middleware.AuthMiddleware())

    // Routes
    r.POST("/products", handlers.ProxyToInventoryService)
    r.GET("/products/:id", handlers.ProxyToInventoryService)
    r.PATCH("/products/:id", handlers.ProxyToInventoryService)
    r.DELETE("/products/:id", handlers.ProxyToInventoryService)
    r.GET("/products", handlers.ProxyToInventoryService)

    r.POST("/orders", handlers.ProxyToOrderService)
    r.GET("/orders/:id", handlers.ProxyToOrderService)
    r.PATCH("/orders/:id", handlers.ProxyToOrderService)
    r.GET("/orders", handlers.ProxyToOrderService)

    r.Run(":8080")
}