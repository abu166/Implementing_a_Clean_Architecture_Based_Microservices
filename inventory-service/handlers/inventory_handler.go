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