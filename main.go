package main

import (
	"embed"
	"go-payments/configs/database"
	"go-payments/internal/payments"
	"go-payments/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//go:embed configs/database/migrations/*.sql
var embedMigrations embed.FS

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Print("error loading .env file, running default settings")
	}

	database.Init_db(embedMigrations)

	utils.Init_custom_validators()

	app := gin.New()

	router := app.Group("api/v1/")
	{
		payments.PaymentRoutes(router)
	}
	app.Run("localhost:8000")
}
