# Inventory Service Documentation

## 1. Project Structure

Here's the folder structure of the inventory-service:

```
inventory-service/
├── config/
│   └── config.go
├── database/
│   └── db.go
├── handlers/
│   └── inventory_handler.go
├── models/
│   └── product.go
├── repositories/
│   └── product_repository.go
├── services/
│   └── inventory_service.go
├── Dockerfile
├── go.mod
├── go.sum
└── main.go
```

This structure adheres to the clean architecture principles:

- **Config**: Handles environment variables and configuration.
- **Database**: Manages database connections and migrations.
- **Handlers**: Defines HTTP endpoints and routes requests to the service layer.
- **Models**: Represents domain models (e.g., Product).
- **Repositories**: Implements data access logic using GORM.
- **Services**: Orchestrates business logic and interacts with repositories.
- **Main**: Initializes the Gin router and starts the server.

## 2. Step-by-Step Explanation

### a. config/config.go

This file manages the application configuration, including database connection details.

#### Code Explanation

```go
package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DbHost     string
    DbPort     string
    DbUser     string
    DbPassword string
    DbName     string
}

var AppConfig *Config

func LoadConfig() {
    err := godotenv.Load("../.env") // Adjust path to point to the root .env file
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    AppConfig = &Config{
        DbHost:     os.Getenv("DB_HOST"),
        DbPort:     os.Getenv("DB_PORT"),
        DbUser:     os.Getenv("DB_USER"),
        DbPassword: os.Getenv("DB_PASSWORD"),
        DbName:     os.Getenv("DB_NAME"),
    }
}
```

**Purpose**:
- Loads environment variables from the .env file.
- Stores configuration in the AppConfig variable.

#### .env File

Ensure your .env file contains the following environment variables:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=abukhassymkhydyrbayev
DB_PASSWORD=admin
DB_NAME=ecommerce
```

### b. database/db.go

This file initializes the GORM database connection and auto-migrates the Product model.

#### Code Explanation

```go
package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/config"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/models"
)

var DB *gorm.DB

func InitDB() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        config.AppConfig.DbHost, config.AppConfig.DbUser, config.AppConfig.DbPassword, config.AppConfig.DbName, config.AppConfig.DbPort)
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }

    // Auto migrate the Product model
    DB.AutoMigrate(&models.Product{})
}
```

**Purpose**:
- Constructs the PostgreSQL DSN (Data Source Name) using configuration values.
- Connects to the database using GORM.
- Automatically migrates the Product model to create the products table.

### c. models/product.go

This file defines the Product model, which represents the products table in the database.

#### Code Explanation

```go
package models

import (
    "time"
)

type Product struct {
    ID          uint    `json:"id" gorm:"primaryKey"`
    Name        string  `json:"name"`
    Category    string  `json:"category"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    Description string  `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

**Purpose**:
- Defines the Product struct with fields corresponding to database columns.
- Tags (json, gorm) map struct fields to JSON and database columns.

### d. repositories/product_repository.go

This file implements CRUD operations for the Product model using GORM.

#### Code Explanation

```go
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
```

**Purpose**:
- Provides repository functions for creating, retrieving, updating, deleting, and listing products.
- Uses GORM to interact with the database.

### e. services/inventory_service.go

This file orchestrates business logic by calling repository functions.

#### Code Explanation

```go
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
```

**Purpose**:
- Acts as an intermediary between the handler layer and the repository layer.
- Calls repository functions to perform CRUD operations.

### f. handlers/inventory_handler.go

This file defines HTTP handlers for the Inventory Service.

#### Code Explanation

```go
package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/models"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/services"
)

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    err := services.CreateProduct(&product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, product)
}

func GetProduct(c *gin.Context) {
    id := c.Param("id")
    product, err := services.GetProductByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var updatedProduct models.Product
    if err := c.ShouldBindJSON(&updatedProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    err := services.UpdateProduct(id, &updatedProduct)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    err := services.DeleteProduct(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func ListProducts(c *gin.Context) {
    products, err := services.ListProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
        return
    }

    c.JSON(http.StatusOK, products)
}
```

**Purpose**:
- Maps HTTP endpoints to service functions.
- Validates input, handles errors, and formats responses.

### g. main.go

This is the entry point of the Inventory Service.

#### Code Explanation

```go
package main

import (
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/config"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/database"
    "Implementing_a_Clean_Architecture_Based_Microservices/inventory-service/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()
    database.InitDB()

    r := gin.Default()

    // Routes
    r.POST("/products", handlers.CreateProduct)
    r.GET("/products/:id", handlers.GetProduct)
    r.PATCH("/products/:id", handlers.UpdateProduct)
    r.DELETE("/products/:id", handlers.DeleteProduct)
    r.GET("/products", handlers.ListProducts)

    r.Run(":8081")
}
```

**Purpose**:
- Loads configuration and initializes the database.
- Sets up the Gin router and registers handlers for each endpoint.
- Starts the server on port 8081.

## 3. Testing the Inventory Service

### Step 1: Start PostgreSQL

Ensure your PostgreSQL database is running and accessible. Use the credentials defined in your .env file.

### Step 2: Build and Run the Inventory Service

Navigate to the inventory-service directory and run:

```bash
docker-compose up --build
```

### Step 3: Test Endpoints

Use curl or Postman to test the Inventory Service.

#### Create a Product

```bash
curl -X POST http://localhost:8081/products \
-H "Content-Type: application/json" \
-d '{
  "name": "Smartphone",
  "category": "Electronics",
  "price": 499.99,
  "stock": 100,
  "description": "A premium smartphone"
}'
```

#### Retrieve a Product

```bash
curl http://localhost:8081/products/1
```

#### Update a Product

```bash
curl -X PATCH http://localhost:8081/products/1 \
-H "Content-Type: application/json" \
-d '{"price": 449.99}'
```

#### Delete a Product

```bash
curl -X DELETE http://localhost:8081/products/1
```

#### List All Products

```bash
curl http://localhost:8081/products
```

## 4. Key Takeaways

- **Configuration**: Environment variables are loaded from .env and used to configure the database connection.
- **Database**: GORM automatically migrates the Product model to create the products table.
- **Handlers**: HTTP endpoints are mapped to service functions, which delegate tasks to repositories.
- **Clean Architecture**: Separation of concerns ensures maintainability and scalability.

## 5. Troubleshooting

If you encounter issues:

### Database Connection Errors:
- Verify the .env file contains correct database credentials.
- Ensure PostgreSQL is running and accessible.

### HTTP Errors:
- Check logs using `docker-compose logs`.
- Ensure all dependencies are installed (`go mod tidy`).