package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/telman03/ecom/db"
	"github.com/telman03/ecom/models"

)
// @Summary Place an order
// @Tags orders
// @Security BearerAuth
// @Success 200 {object} map[string]string "message"
// @Router /orders/place [post]
func PlaceOrder(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Fetch cart items
	var cartItems []models.Cart
	db.DB.Where("user_id = ?", userID).Find(&cartItems)

	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	// Calculate total price
	var total float64
	var orderItems []models.OrderItem
	for _, item := range cartItems {
		var product models.Product
		db.DB.First(&product, item.ProductID)

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
		total += product.Price * float64(item.Quantity)
	}

	// Create order
	order := models.Order{UserID: userID.(uint), Items: orderItems, Total: total}
	db.DB.Create(&order)

	// Clear user's cart
	db.DB.Where("user_id = ?", userID).Delete(&models.Cart{})

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully"})
}

// @Summary Get all orders
// @Tags orders
// @Security BearerAuth
// @Success 200 {array} models.ResponseMessage
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var orders []models.Order
	db.DB.Preload("Items").Where("user_id = ?", userID).Find(&orders)

	c.JSON(http.StatusOK, orders)
}

// @Summary Get order details
// @Tags orders
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {array} models.ResponseMessage
// @Router /orders/{id} [get]
func GetOrderDetails(c *gin.Context) {
	orderID := c.Param("id")

	var order models.Order
	if err := db.DB.Preload("Items").First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
