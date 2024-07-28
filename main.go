package main

import (
	"embed"
	"go-payments/configs/database"
	"go-payments/internal/expenses/api"
	"go-payments/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//go:embed configs/database/migrations/*.sql
var embed_migrations embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("error loading .env file, running default settings")
	}
	database.InitDb(embed_migrations)
	utils.InitCustomValidators()
	app := gin.New()
	router := app.Group("api/v1/")
	{
		api.ExpensesRoutes(router)
	}
	app.Run("localhost:8000")
}
