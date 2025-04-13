package repositories

import (
    "errors"
    "gorm.io/gorm"
    "Implementing_a_Clean_Architecture_Based_Microservices/user-service/database"
    "Implementing_a_Clean_Architecture_Based_Microservices/user-service/models"
)

func CreateUser(user *models.User) error {
    return database.DB.Create(user).Error
}

func GetUserByUsername(username string) (*models.User, error) {
    user := &models.User{}
    err := database.DB.Where("username = ?", username).First(user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errors.New("User not found")
    }
    return user, err
}

func GetUserByID(id int32) (*models.User, error) {
    user := &models.User{}
    err := database.DB.First(user, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errors.New("User not found")
    }
    return user, err
}