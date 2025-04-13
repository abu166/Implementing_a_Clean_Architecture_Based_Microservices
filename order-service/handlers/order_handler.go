package handlers

import (
	"Implementing_a_Clean_Architecture_Based_Microservices/order-service/models"
	"Implementing_a_Clean_Architecture_Based_Microservices/order-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Helper function to convert string to int32 safely
func convertToInt32(value string) (int32, error) {
	num, err := strconv.Atoi(value)
	if err != nil || num < 0 {
		return 0, err // Return an error if the value is invalid or negative
	}
	return int32(num), nil
}

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
	productID, err := convertToInt32(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := services.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Bind the request body to the product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Convert the ID parameter to int32
	productID, err := convertToInt32(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Set the ID field of the product struct
	product.ID = uint(productID)

	// Call the service layer function
	err = services.UpdateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated product
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := convertToInt32(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = services.DeleteProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func ListProducts(c *gin.Context) {
	category := c.Query("category")
	page := convertToInt32Default(c.DefaultQuery("page", "1"))
	limit := convertToInt32Default(c.DefaultQuery("limit", "10"))

	products, err := services.ListProducts(category, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Helper function to convert string to int32 with default value
func convertToInt32Default(value string) int32 {
	num, _ := strconv.Atoi(value)
	return int32(num)
}
