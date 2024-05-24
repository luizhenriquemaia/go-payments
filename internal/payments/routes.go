package payments

import "github.com/gin-gonic/gin"

func getPaymentsRoute(c *gin.Context) {
	data := getPaymentsController()
	c.IndentedJSON(200, gin.H{"data": data})
}
