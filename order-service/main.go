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
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products/:id", handlers.GetProduct)
	r.PATCH("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
	r.GET("/products", handlers.ListProducts)

	r.Run(":8082")
}
