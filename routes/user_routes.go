package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/controllers"
	"github.com/telman03/ecom/middleware"
)

func UserRoutes(r *gin.Engine) {
	protected := r.Group("/user")
	protected.Use(middleware.AuthMiddleware()) // Protect these routes

	protected.GET("/profile", controllers.GetProfile)
	protected.GET("/orders", controllers.GetOrders)
}
