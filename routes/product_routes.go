package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/controllers"
	"github.com/telman03/ecom/middleware"
)

func ProductRoutes(r *gin.Engine) {
	// Public Routes
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)

	// Protected Routes (Require Auth)
	protected := r.Group("/products")
	protected.Use(middleware.AuthMiddleware())

	protected.POST("/", controllers.CreateProduct)
	protected.PUT("/:id", controllers.UpdateProduct)
	protected.DELETE("/:id", controllers.DeleteProduct)
}
