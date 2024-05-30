package payments

import "github.com/gin-gonic/gin"

func getPaymentsRoute(c *gin.Context) {
	data := getPaymentsController()
	c.IndentedJSON(200, gin.H{"data": data})
}

func postPaymentsRoute(c *gin.Context) {
	entity, err := addPaymentController(c)
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.IndentedJSON(201, gin.H{"data": entity})
}
