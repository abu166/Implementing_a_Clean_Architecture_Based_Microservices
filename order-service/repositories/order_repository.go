package repositories

import (
    "errors"
	"gorm.io/gorm"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/database"
    "Implementing_a_Clean_Architecture_Based_Microservices/order-service/models"
)

func CreateProduct(product *models.Product) error {
    return database.DB.Create(product).Error
}

func GetProductByID(id int32) (*models.Product, error) {
    product := &models.Product{}
    err := database.DB.First(product, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errors.New("Product not found")
    }
    return product, err
}

func UpdateProduct(product *models.Product) error {
    return database.DB.Save(product).Error
}

func DeleteProduct(id int32) error {
    return database.DB.Delete(&models.Product{}, id).Error
}

func ListProducts(category string, page, limit int32) ([]models.Product, error) {
    var products []models.Product
    query := database.DB
    if category != "" {
        query = query.Where("category = ?", category)
    }
    if page > 0 && limit > 0 {
        offset := (page - 1) * limit
        query = query.Offset(int(offset)).Limit(int(limit))
    }
    err := query.Find(&products).Error
    return products, err
}