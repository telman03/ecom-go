package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/telman03/ecom/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/telman03/ecom/db"
	"github.com/telman03/ecom/models"
	"github.com/telman03/ecom/routes"
)
// @title E-commerce API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db.InitDB()

	db.DB.AutoMigrate(&models.User{}, &models.Order{}, &models.OrderItem{}, &models.Product{}, &models.Cart{})

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.ProductRoutes(r)
    routes.RegisterCartRoutes(r)
    routes.RegisterOrderRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
