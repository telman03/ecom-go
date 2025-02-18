package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/db"
	"github.com/telman03/ecom/models"
)

// Add Product to Cart
// @Summary Add a product to the cart
// @Description Add a specific quantity of a product to the user's cart
// @Accept json
// @Tags cart
// @Produce json
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /cart [post]
func AddToCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input models.Cart
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the product exists
	var product models.Product
	if err := db.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Add or update cart item
	var cartItem models.Cart
	if err := db.DB.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&cartItem).Error; err != nil {
		// Product not found in cart, create a new entry
		input.UserID = userID.(uint)
		if err := db.DB.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to cart"})
			return
		}
	} else {
		// Update existing cart item
		cartItem.Quantity += input.Quantity
		db.DB.Save(&cartItem)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}

// Get Cart
// @Summary Get the user's cart
// @Description Get all items in the user's cart
// @Produce json
// @Tags cart
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /cart [get]
func GetCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var cartItems []models.Cart
	if err := db.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}


// Remove Product from Cart
// @Summary Remove a product from the cart
// @Tags cart
// @Description Remove a specific product from the user's cart by its ID
// @Param id path int true "Product ID"
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /cart/{id} [delete]
func RemoveFromCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	productID := c.Param("id")
	if err := db.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Cart{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove product from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
}
