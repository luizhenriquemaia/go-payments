package main

import (
	"embed"
	"go-payments/configs/db"
	"go-payments/internal/payments"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//go:embed configs/db/migrations/*.sql
var embedMigrations embed.FS

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Print("error loading .env file, running default settings")
	}

	db.Init_db(embedMigrations)

	app := gin.New()

	router := app.Group("api/v1/")
	{
		payments.PaymentRoutes(router)
	}
	app.Run("localhost:8000")
}
