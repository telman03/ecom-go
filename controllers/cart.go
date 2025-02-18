package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/db"
	"github.com/telman03/ecom/models"
)

// @Summary Add item to cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body map[string]int true "Product ID and Quantity"
// @Success 200 {object} map[string]string "message"
// @Failure 400 {object} map[string]string "error"
// @Router /cart/add [post]
func AddToCart(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartItem := models.Cart{UserID: userID.(uint), ProductID: input.ProductID, Quantity: input.Quantity}
	db.DB.Create(&cartItem)

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

// @Summary View cart
// @Tags cart
// @Security BearerAuth
// @Success 200 {array} models.ResponseMessage
// @Router /cart [get]
func GetCart(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var cart []models.Cart
	db.DB.Where("user_id = ?", userID).Find(&cart)

	c.JSON(http.StatusOK, cart)
}

// @Summary Update cart item
// @Tags cart
// @Accept json
// @Security BearerAuth
// @Param data body map[string]int true "Cart ID and Quantity"
// @Success 200 {object} map[string]string "message"
// @Router /cart/update [put]
func UpdateCart(c *gin.Context) {
	var input struct {
		CartID   uint `json:"cart_id"`
		Quantity int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&models.Cart{}).Where("id = ?", input.CartID).Update("quantity", input.Quantity)
	c.JSON(http.StatusOK, gin.H{"message": "Cart updated"})
}

// @Summary Remove item from cart
// @Tags cart
// @Security BearerAuth
// @Param cart_id query int true "Cart ID"
// @Success 200 {object} map[string]string "message"
// @Router /cart/remove [delete]
func RemoveFromCart(c *gin.Context) {
	cartID := c.Query("cart_id")
	db.DB.Delete(&models.Cart{}, cartID)

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
