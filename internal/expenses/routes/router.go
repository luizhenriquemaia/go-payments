package routes

import "github.com/gin-gonic/gin"

func ExpensesRoutes(superRouter *gin.RouterGroup) {
	router := superRouter.Group("/expense")
	{
		router.GET("", getExpensesRoute)
		router.POST("/", postExpensesRoute)
		router.POST("pay/", postPaymentRoute)
	}
}
