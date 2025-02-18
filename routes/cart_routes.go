package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/controllers"
	"github.com/telman03/ecom/middleware"
)

func CartRoutes(r *gin.Engine) {
	protected := r.Group("/cart")
	protected.Use(middleware.AuthMiddleware()) // Protect cart routes

	protected.POST("/", controllers.AddToCart)
	protected.GET("/", controllers.GetCart)
	protected.DELETE("/:id", controllers.RemoveFromCart)
}

func OrderRoutes(r *gin.Engine) {
	protected := r.Group("/order")
	protected.Use(middleware.AuthMiddleware()) // Protect order routes

	protected.POST("/", controllers.PlaceOrder)
}
