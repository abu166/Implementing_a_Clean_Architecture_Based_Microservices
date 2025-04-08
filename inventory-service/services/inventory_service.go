package services

import (
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/models"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/repositories"
)

func CreateProduct(product *models.Product) error {
    return repositories.CreateProduct(product)
}

func GetProductByID(id string) (*models.Product, error) {
    return repositories.GetProductByID(id)
}

func UpdateProduct(id string, product *models.Product) error {
    return repositories.UpdateProduct(id, product)
}

func DeleteProduct(id string) error {
    return repositories.DeleteProduct(id)
}

func ListProducts() ([]models.Product, error) {
    return repositories.ListProducts()
}