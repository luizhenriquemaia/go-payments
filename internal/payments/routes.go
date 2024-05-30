package payments

import (
	"errors"
	"go-payments/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func getPaymentsRoute(c *gin.Context) {
	data := getPaymentsController()
	c.IndentedJSON(200, gin.H{"data": data})
}

func postPaymentsRoute(c *gin.Context) {
	entity, err := addPaymentController(c)
	if err != nil {
		var validation_errors validator.ValidationErrors
		if errors.As(err, &validation_errors) {
			api_error := utils.Get_validation_api_error(validation_errors)
			c.IndentedJSON(400, gin.H{"msg": err.Error(), "errors": api_error})
			return
		}
		c.IndentedJSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.IndentedJSON(201, gin.H{"data": entity})
}
