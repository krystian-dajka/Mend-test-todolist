package main

import (
	"context"
	"os"
	"time"
	"log"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/krystian-dajka/Mend-test-todolist/config"
	"github.com/krystian-dajka/Mend-test-todolist/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := config.ConnectDB()

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// basic security headers
	router.Use(helmet.Default())

	routes.SetupRouter(router, client)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000" // Default port
	}

	// Start server on specified port
	router.Run(":" + port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)

}
