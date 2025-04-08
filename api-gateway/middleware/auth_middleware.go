package middleware

import (
    "github.com/gin-gonic/gin"
    "strings"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // List of public routes (no authentication required)
        publicRoutes := map[string]bool{
            "GET /products/:id": true,
            "GET /orders":       true,
            "GET /orders/:id":   true,
        }

        // Check if the current route is public
        key := c.Request.Method + " " + c.FullPath()
        if publicRoutes[key] {
            c.Next()
            return
        }

        // Enforce authentication for other routes
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Missing Authorization header"})
            c.Abort()
            return
        }

        // Extract and validate the token (optional for now)
        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        if tokenString == "" {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // If the token is valid, proceed to the next handler
        c.Next()
    }
}