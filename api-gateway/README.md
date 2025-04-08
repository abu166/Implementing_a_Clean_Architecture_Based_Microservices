# API Gateway Documentation

## 1. Overview of the API Gateway

The API Gateway acts as a central entry point for all client requests. It:

- **Routes Requests**: Proxies incoming requests to the Inventory Service or Order Service based on the endpoint.
- **Applies Middleware**: Adds functionality like logging, authentication, and error handling before forwarding requests.
- **Handles Responses**: Forwards responses from the downstream services back to the client.

## 2. Key Components

### a. config/config.go

This file manages configuration settings for the API Gateway, including the URLs of the Inventory Service and Order Service.

#### Code Explanation

```go
package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    InventoryServiceURL string
    OrderServiceURL     string
}

var AppConfig *Config

func LoadConfig() {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    AppConfig = &Config{
        InventoryServiceURL: os.Getenv("INVENTORY_SERVICE_URL"),
        OrderServiceURL:     os.Getenv("ORDER_SERVICE_URL"),
    }
}
```

- **Purpose**: Loads environment variables from the .env file.
- **Environment Variables**:
  - `INVENTORY_SERVICE_URL`: URL of the Inventory Service (e.g., http://localhost:8081).
  - `ORDER_SERVICE_URL`: URL of the Order Service (e.g., http://localhost:8082).

### b. handlers/gateway_handler.go

This file contains handlers that proxy requests to the Inventory Service and Order Service.

#### Code Explanation

```go
package handlers

import (
    "bytes"
    "io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
    "Implementing_a_Clean_Architecture_Based_Microservices/api-gateway/config"
)

// ProxyToInventoryService proxies requests to the Inventory Service
func ProxyToInventoryService(c *gin.Context) {
    proxyRequest(c, Config().InventoryServiceURL)
}

// ProxyToOrderService proxies requests to the Order Service
func ProxyToOrderService(c *gin.Context) {
    proxyRequest(c, Config().OrderServiceURL)
}

func proxyRequest(c *gin.Context, targetURL string) {
    method := c.Request.Method
    url := targetURL + c.Request.URL.Path

    reqBody, _ := ioutil.ReadAll(c.Request.Body)
    req, _ := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
    req.Header = c.Request.Header

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to proxy request"})
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

// Config retrieves the application configuration
func Config() *config.Config {
    return config.AppConfig
}
```

#### Proxy Logic:
1. **Extract Request Details**:
   - `method`: HTTP method of the incoming request (POST, GET, etc.).
   - `url`: Full URL of the target service (e.g., http://localhost:8081/products).
   - `reqBody`: Body of the incoming request.

2. **Create New Request**:
   - Constructs a new HTTP request with the same method, URL, headers, and body.

3. **Send Request to Target Service**:
   - Uses an `http.Client` to send the request to the Inventory Service or Order Service.

4. **Handle Response**:
   - Reads the response body and forwards it back to the client with the same status code and content type.

### c. middleware/auth_middleware.go

This file defines the authentication middleware, which enforces token-based authentication for certain routes.

#### Code Explanation

```go
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
```

#### Authentication Logic:
1. **Public Routes**:
   - Certain routes (e.g., GET /products/:id, GET /orders) are marked as public and do not require authentication.

2. **Authorization Header**:
   - Checks for the presence of an Authorization header.
   - Validates the token (currently placeholder logic; can be extended for JWT validation).

3. **Forward Request**:
   - If the token is valid, allows the request to proceed to the next handler.

### d. middleware/logging_middleware.go

This file defines the logging middleware, which logs details about incoming requests.

#### Code Explanation

```go
package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        // Process the request
        c.Next()

        // Calculate duration
        end := time.Since(start)

        // Log the request details
        log.Printf(
            "%s %s %d %s %v",
            c.Request.Method,
            c.Request.URL.Path,
            c.Writer.Status(),
            c.ClientIP(),
            end,
        )
    }
}
```

#### Logging Logic:
1. **Start Time**: Records the start time of the request.
2. **Process Request**: Allows the request to proceed through the pipeline.
3. **End Time**: Records the end time of the request.
4. **Log Details**:
   - Method (GET, POST, etc.).
   - Path (e.g., /products, /orders).
   - Status Code (e.g., 200, 404).
   - Client IP.
   - Duration of the request.

### e. main.go

This is the entry point of the API Gateway. It initializes Gin, registers middleware, defines routes, and starts the server.

#### Code Explanation

```go
package main

import (
    "github.com/gin-gonic/gin"
    "Implementing_a_Clean_Architecture_Based_Microservices/api-gateway/handlers"
    "Implementing_a_Clean_Architecture_Based_Microservices/api-gateway/middleware"
    "Implementing_a_Clean_Architecture_Based_Microservices/api-gateway/config"
)

func main() {
    config.LoadConfig()

    r := gin.Default()

    // Middleware
    r.Use(middleware.LoggingMiddleware())
    r.Use(middleware.AuthMiddleware())

    // Routes
    r.POST("/products", handlers.ProxyToInventoryService)
    r.GET("/products/:id", handlers.ProxyToInventoryService)
    r.PATCH("/products/:id", handlers.ProxyToInventoryService)
    r.DELETE("/products/:id", handlers.ProxyToInventoryService)
    r.GET("/products", handlers.ProxyToInventoryService)

    r.POST("/orders", handlers.ProxyToOrderService)
    r.GET("/orders/:id", handlers.ProxyToOrderService)
    r.PATCH("/orders/:id", handlers.ProxyToOrderService)
    r.GET("/orders", handlers.ProxyToOrderService)

    r.Run(":8080")
}
```

#### Initialization:
1. **Load Configuration**: Loads environment variables using `config.LoadConfig()`.
2. **Initialize Gin**: Creates a new Gin router.
3. **Register Middleware**:
   - `LoggingMiddleware`: Logs request details.
   - `AuthMiddleware`: Enforces authentication.
4. **Define Routes**:
   - Maps endpoints to their respective handlers (`ProxyToInventoryService` and `ProxyToOrderService`).
5. **Start Server**: Runs the server on port 8080.

## 3. How It Works Step by Step

### Step 1: Client Sends a Request

A client sends a request to the API Gateway (e.g., POST http://localhost:8080/products).

### Step 2: Middleware Executes

1. **Logging Middleware**:
   - Logs the request details (method, path, status code, client IP, duration).

2. **Authentication Middleware**:
   - Checks if the route requires authentication.
   - Validates the Authorization header (if required).

### Step 3: Route Matching

The API Gateway matches the incoming request to the appropriate handler:
- POST /products → ProxyToInventoryService
- POST /orders → ProxyToOrderService

### Step 4: Proxy Request

The proxyRequest function constructs a new HTTP request and forwards it to the target service:

1. **Extract Request Details**:
   - Method (POST), URL (http://localhost:8081/products), Headers, Body.

2. **Send Request**:
   - Uses http.Client to send the request to the Inventory Service or Order Service.

3. **Handle Response**:
   - Reads the response from the target service and forwards it back to the client.

### Step 5: Client Receives Response

The client receives the response from the target service via the API Gateway.

## 4. Example Workflow

### Scenario: Creating a Product

1. **Client Sends Request**:
```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-H "Authorization: Bearer dummy-token" \
-d '{
  "name": "Smartphone",
  "category": "Electronics",
  "price": 499.99,
  "stock": 100,
  "description": "A premium smartphone"
}'
```

2. **API Gateway Processes Request**:
   - Logging Middleware: Logs the request details.
   - Authentication Middleware: Validates the Authorization header.
   - Route Matching: Matches the request to ProxyToInventoryService.

3. **Proxy Request**:
   - Forwards the request to http://localhost:8081/products.

4. **Inventory Service Processes Request**:
   - Creates the product in the database.
   - Returns a 201 Created response with the product details.

5. **API Gateway Forwards Response**:
   - Passes the response back to the client.

6. **Client Receives Response**:
```json
{
  "id": 1,
  "name": "Smartphone",
  "category": "Electronics",
  "price": 499.99,
  "stock": 100,
  "description": "A premium smartphone",
  "created_at": "2025-04-07T15:51:05Z",
  "updated_at": "2025-04-07T15:51:05Z"
}
```

## 5. Key Takeaways

- **Routing**: The API Gateway routes requests to the appropriate microservices.
- **Middleware**: Enhances functionality with logging, authentication, and error handling.
- **Proxying**: Forwards requests to downstream services and handles responses.