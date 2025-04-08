package models

import (
    "time"
)

type Order struct {
    ID         uint      `json:"id" gorm:"primaryKey"`
    UserID     uint      `json:"user_id"`
    Status     string    `json:"status"`
    TotalPrice float64   `json:"total_price"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
    ID       uint    `json:"id" gorm:"primaryKey"`
    OrderID  uint    `json:"order_id"`
    Product  string  `json:"product"`
    Quantity int     `json:"quantity"`
    Price    float64 `json:"price"`
}