package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/controllers"
	"github.com/telman03/ecom/middleware"
)

func RegisterCartRoutes(r *gin.Engine) {
	cartGroup := r.Group("/cart")
	cartGroup.Use(middleware.AuthMiddleware()) // Protect routes

	cartGroup.POST("/add", controllers.AddToCart)
	cartGroup.GET("/", controllers.GetCart)
	cartGroup.PUT("/update", controllers.UpdateCart)
	cartGroup.DELETE("/remove", controllers.RemoveFromCart)
}
