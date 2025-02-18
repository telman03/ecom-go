package main

import (
	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/db"
	"github.com/telman03/ecom/models"
	"github.com/telman03/ecom/routes"
)

func main() {
    db.InitDB()


    db.DB.AutoMigrate(&models.User{}, &models.Order{}, &models.Product{})

    r := gin.Default()

    routes.AuthRoutes(r)
    routes.UserRoutes(r) 
    routes.ProductRoutes(r)

    r.Run(":8080")
}
