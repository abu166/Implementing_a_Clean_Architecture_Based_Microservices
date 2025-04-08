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