package main

import (
	"go-payments/internal/payments"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	router := app.Group("api/v1/")
	{
		payments.PaymentRoutes(router)
	}

	app.Run("localhost:8000")
}
