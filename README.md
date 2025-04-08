# Implementing_a_Clean_Architecture_Based_Microservices

# Microservices API Test Cases

## 1. Inventory Service Test Cases

### 1.1 Create a Product

This command creates a new product in the Inventory Service.

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

Expected Response:

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

### 1.2 Retrieve a Product by ID

This command retrieves a product by its ID.

```bash
curl http://localhost:8081/products/1
```

Expected Response:

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

### 1.3 Update a Product

This command updates the price of an existing product.

```bash
curl -X PATCH http://localhost:8081/products/1 \
-H "Content-Type: application/json" \
-d '{"price": 449.99}'
```

Expected Response:

```json
{
  "message": "Product updated"
}
```

### 1.4 Delete a Product

This command deletes a product by its ID.

```bash
curl -X DELETE http://localhost:8081/products/1
```

Expected Response:

```json
{
  "message": "Product deleted"
}
```

## 2. Order Service Test Cases

### 2.1 Create an Order

This command creates a new order in the Order Service.

```bash
curl -X POST http://localhost:8082/orders \
-H "Content-Type: application/json" \
-d '{
  "user_id": 1,
  "status": "pending",
  "total_price": 1999.98,
  "items": [
    {
      "product": "Laptop",
      "quantity": 2,
      "price": 999.99
    }
  ]
}'
```

Expected Response:

```json
{
  "id": 1,
  "user_id": 1,
  "status": "pending",
  "total_price": 1999.98,
  "created_at": "2025-04-07T15:55:00Z",
  "updated_at": "2025-04-07T15:55:00Z",
  "items": [
    {
      "id": 1,
      "order_id": 1,
      "product": "Laptop",
      "quantity": 2,
      "price": 999.99
    }
  ]
}
```

### 2.2 Retrieve an Order by ID

This command retrieves an order by its ID.

```bash
curl http://localhost:8082/orders/1
```

Expected Response:

```json
{
  "id": 1,
  "user_id": 1,
  "status": "pending",
  "total_price": 1999.98,
  "created_at": "2025-04-07T15:55:00Z",
  "updated_at": "2025-04-07T15:55:00Z",
  "items": [
    {
      "id": 1,
      "order_id": 1,
      "product": "Laptop",
      "quantity": 2,
      "price": 999.99
    }
  ]
}
```

### 2.3 Update Order Status

This command updates the status of an existing order.

```bash
curl -X PATCH http://localhost:8082/orders/1 \
-H "Content-Type: application/json" \
-d '{"status": "completed"}'
```

Expected Response:

```json
{
  "message": "Order status updated"
}
```

### 2.4 List All Orders

This command retrieves all orders.

```bash
curl http://localhost:8082/orders
```

Expected Response:

```json
[
  {
    "id": 1,
    "user_id": 1,
    "status": "completed",
    "total_price": 1999.98,
    "created_at": "2025-04-07T15:55:00Z",
    "updated_at": "2025-04-07T15:56:00Z",
    "items": [
      {
        "id": 1,
        "order_id": 1,
        "product": "Laptop",
        "quantity": 2,
        "price": 999.99
      }
    ]
  }
]
```

## 3. API Gateway Test Cases

### 3.1 Create a Product via API Gateway

This command creates a new product through the API Gateway.

```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-H "Authorization: Bearer dummy-token" \
-d '{
  "name": "Tablet",
  "category": "Electronics",
  "price": 299.99,
  "stock": 30,
  "description": "A lightweight tablet"
}'
```

Expected Response:

```json
{
  "id": 2,
  "name": "Tablet",
  "category": "Electronics",
  "price": 299.99,
  "stock": 30,
  "description": "A lightweight tablet",
  "created_at": "2025-04-07T16:00:00Z",
  "updated_at": "2025-04-07T16:00:00Z"
}
```

### 3.2 Retrieve a Product via API Gateway

This command retrieves a product through the API Gateway.

```bash
curl http://localhost:8080/products/2 \
-H "Authorization: Bearer dummy-token"
```

Expected Response:

```json
{
  "id": 2,
  "name": "Tablet",
  "category": "Electronics",
  "price": 299.99,
  "stock": 30,
  "description": "A lightweight tablet",
  "created_at": "2025-04-07T16:00:00Z",
  "updated_at": "2025-04-07T16:00:00Z"
}
```

### 3.3 Create an Order via API Gateway

This command creates a new order through the API Gateway.

```bash
curl -X POST http://localhost:8080/orders \
-H "Content-Type: application/json" \
-H "Authorization: Bearer dummy-token" \
-d '{
  "user_id": 1,
  "status": "pending",
  "total_price": 1999.98,
  "items": [
    {
      "product": "Laptop",
      "quantity": 2,
      "price": 999.99
    }
  ]
}'
```

Expected Response:

```json
{
  "id": 3,
  "user_id": 1,
  "status": "pending",
  "total_price": 1999.98,
  "created_at": "2025-04-07T16:05:00Z",
  "updated_at": "2025-04-07T16:05:00Z",
  "items": [
    {
      "id": 3,
      "order_id": 3,
      "product": "Laptop",
      "quantity": 2,
      "price": 999.99
    }
  ]
}
```