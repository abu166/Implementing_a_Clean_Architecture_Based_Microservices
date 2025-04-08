package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/models"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/services"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    err := services.CreateOrder(&order)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, order)
}

func GetOrder(c *gin.Context) {
    id := c.Param("id")
    order, err := services.GetOrderByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}

func UpdateOrderStatus(c *gin.Context) {
    id := c.Param("id")
    var statusUpdate struct {
        Status string `json:"status"`
    }
    if err := c.ShouldBindJSON(&statusUpdate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    err := services.UpdateOrderStatus(id, statusUpdate.Status)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found or invalid status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order status updated"})
}

func ListOrders(c *gin.Context) {
    orders, err := services.ListOrders()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, orders)
}