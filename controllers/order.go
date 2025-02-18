package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/db"
	"github.com/telman03/ecom/models"
)

// Get Orders
// @Summary Get a list of orders for the logged-in user
// @Description Fetches all the orders associated with the authenticated user
// @Accept json
// @Produce json
// @Tags order
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var orders []models.Order
	if err := db.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// Place Order
// @Summary Place an order
// @Description Place an order for all items in the user's cart
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /order [post]
func PlaceOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user's cart
	var cartItems []models.Cart
	if err := db.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}

	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	// Calculate the total price
	var totalPrice float64
	var orderItems []models.OrderItem
	for _, item := range cartItems {
		var product models.Product
		db.DB.First(&product, item.ProductID)
		totalPrice += product.Price * float64(item.Quantity)
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
	}

	// Create order
	order := models.Order{
		UserID:     userID.(uint),
		TotalPrice: totalPrice,
		Status:     "pending",
		Items:      orderItems,
	}
	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}

	// Remove items from cart
	db.DB.Where("user_id = ?", userID).Delete(&models.Cart{})

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order_id": order.ID})
}