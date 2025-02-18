package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint     `json:"user_id"`
	Items  []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Total  float64  `json:"total"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint   `json:"order_id"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     float64 `json:"price"`
}
