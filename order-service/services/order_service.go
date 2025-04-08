package services

import (
    "errors"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/models"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/repositories"
)

func CreateOrder(order *models.Order) error {
    return repositories.CreateOrder(order)
}

func GetOrderByID(id string) (*models.Order, error) {
    return repositories.GetOrderByID(id)
}

func UpdateOrderStatus(id string, status string) error {
    validStatuses := map[string]bool{
        "pending":   true,
        "completed": true,
        "cancelled": true,
    }

    if !validStatuses[status] {
        return errors.New("Invalid order status")
    }

    return repositories.UpdateOrderStatus(id, status)
}

func ListOrders() ([]models.Order, error) {
    return repositories.ListOrders()
}