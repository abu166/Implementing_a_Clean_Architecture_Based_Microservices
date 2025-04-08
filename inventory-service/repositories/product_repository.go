package repositories

import (
    "errors"
	"gorm.io/gorm"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/database"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/models"
)

func CreateProduct(product *models.Product) error {
    return database.DB.Create(product).Error
}

func GetProductByID(id string) (*models.Product, error) {
    product := &models.Product{}
    err := database.DB.First(product, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errors.New("Product not found")
    }
    return product, err
}

func UpdateProduct(id string, product *models.Product) error {
    return database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func DeleteProduct(id string) error {
    return database.DB.Delete(&models.Product{}, id).Error
}

func ListProducts() ([]models.Product, error) {
    var products []models.Product
    err := database.DB.Find(&products).Error
    return products, err
}