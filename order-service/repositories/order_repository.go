package repositories

import (
    "errors"
	"gorm.io/gorm"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/database"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/models"
)

func CreateOrder(order *models.Order) error {
    return database.DB.Create(order).Error
}

func GetOrderByID(id string) (*models.Order, error) {
    order := &models.Order{}
    err := database.DB.Preload("Items").First(order, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errors.New("Order not found")
    }
    return order, err
}

func UpdateOrderStatus(id string, status string) error {
    return database.DB.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}

func ListOrders() ([]models.Order, error) {
    var orders []models.Order
    err := database.DB.Preload("Items").Find(&orders).Error
    return orders, err
}