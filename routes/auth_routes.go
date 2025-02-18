package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/controllers"
)


func AuthRoutes(r *gin.Engine){
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}