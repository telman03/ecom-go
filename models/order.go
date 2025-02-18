package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint
	TotalPrice float64
	Status     string `gorm:"default:'pending'"`
	Items      []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
}
