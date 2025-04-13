package services

import (
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/models"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/repositories"
)

func CreateProduct(product *models.Product) error {
    return repositories.CreateProduct(product)
}

func GetProduct(id int32) (*models.Product, error) {
    return repositories.GetProductByID(id)
}

func UpdateProduct(product *models.Product) error {
    return repositories.UpdateProduct(product)
}

func DeleteProduct(id int32) error {
    return repositories.DeleteProduct(id)
}

func ListProducts(category string, page, limit int32) ([]models.Product, error) {
    return repositories.ListProducts(category, page, limit)
}