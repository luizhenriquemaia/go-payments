package payments

import "github.com/gin-gonic/gin"

func PaymentRoutes(superRouter *gin.RouterGroup) {
	router := superRouter.Group("/payment")
	{
		router.GET("", getPaymentsRoute)
		router.POST("/", postPaymentsRoute)
	}
}
