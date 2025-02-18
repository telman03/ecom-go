package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/db"
	"github.com/telman03/ecom/models"
)


// Get Profile
// @Summary Get the user's profile information
// @Description Fetches the details of the logged-in user (username, email) based on the provided JWT token
// @Accept json
// @Tags auth
// @SecurityBearer
// @Produce json
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /user/profile [get]
func GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}