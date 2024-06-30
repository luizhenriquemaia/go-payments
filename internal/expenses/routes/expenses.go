package routes

import (
	"errors"
	"go-payments/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go-payments/internal/expenses/controllers"
)

func getExpensesRoute(c *gin.Context) {
	entitites, err := controllers.GetExpensesController(c)
	if err != nil {
		c.IndentedJSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"data": entitites})
}

func postExpensesRoute(c *gin.Context) {
	entity, err := controllers.AddExpensesController(c)
	if err != nil {
		var validation_errors validator.ValidationErrors
		if errors.As(err, &validation_errors) {
			api_error := utils.Get_validation_api_error(validation_errors)
			c.IndentedJSON(400, gin.H{"msg": "Invalid data", "errors": api_error})
			return
		}
		c.IndentedJSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.IndentedJSON(201, gin.H{"data": entity})
}
