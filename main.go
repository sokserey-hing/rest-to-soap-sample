package main

import (
    "gin-soap-service/handlers"
    "gin-soap-service/utils"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    // Load environment variables
    if err := utils.LoadEnv(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	PORT := utils.GetEnv("PORT", "8080")

    // Set up Gin router
    router := gin.Default()

    // Define route
    router.POST("/convert", handlers.ConvertHandler)

    // Start server
    router.Run(":" + PORT)
}
