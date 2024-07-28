package api

import (
	pay_api "go-payments/internal/payments/api"

	"github.com/gin-gonic/gin"
)

func ExpensesRoutes(superRouter *gin.RouterGroup) {
	router := superRouter.Group("/expense")
	{
		router.GET("", getExpensesRoute)
		router.POST("/", postExpensesRoute)
		router.POST("pay/", pay_api.PostPaymentRoute)
	}
}
