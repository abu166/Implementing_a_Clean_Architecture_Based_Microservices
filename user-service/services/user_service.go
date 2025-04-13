package services

import (
    "errors"
    "Implementing_a_Clean_Architecture_Based_Microservices/user-service/models"
    "Implementing_a_Clean_Architecture_Based_Microservices/user-service/repositories"
)

func RegisterUser(username, password, email string) (*models.User, error) {
    user := &models.User{
        Username: username,
        Password: password,
        Email:    email,
    }

    err := repositories.CreateUser(user)
    if err != nil {
        return nil, errors.New("Failed to register user")
    }

    return user, nil
}

func AuthenticateUser(username, password string) (string, error) {
    user, err := repositories.GetUserByUsername(username)
    if err != nil || user.Password != password {
        return "", errors.New("Invalid credentials")
    }

    // Generate a simple token (replace with JWT if needed)
    return "token-" + user.Username, nil
}

func GetUserProfile(id int32) (*models.User, error) {
    return repositories.GetUserByID(id)
}