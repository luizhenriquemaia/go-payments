package routes

import "github.com/gin-gonic/gin"

func ExpensesRoutes(superRouter *gin.RouterGroup) {
	router := superRouter.Group("/payment")
	{
		router.GET("", getExpensesRoute)
		router.POST("/", postExpensesRoute)
	}
}
