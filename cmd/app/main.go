package main

import (
	"go-payments/configs"
	"go-payments/internal/payments"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file, running default settings")
	}

	configs.Init_db()
	app := gin.New()
	router := app.Group("api/v1/")
	{
		payments.PaymentRoutes(router)
	}
	app.Run("localhost:8000")
}
