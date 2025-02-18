package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/controllers"
	"github.com/telman03/ecom/middleware"
)

func RegisterOrderRoutes(r *gin.Engine) {
	orderGroup := r.Group("/orders")
	orderGroup.Use(middleware.AuthMiddleware()) // Protect order routes

	orderGroup.POST("/place", controllers.PlaceOrder)
	orderGroup.GET("/", controllers.GetOrders)
	orderGroup.GET("/:id", controllers.GetOrderDetails)
}
