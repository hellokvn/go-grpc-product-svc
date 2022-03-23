package models

import "gorm.io/gorm"

type StockDecreaseLog struct {
	gorm.Model
	Id      int `json:"id" gorm:"primaryKey"`
	OrderId int `json:"order_id"`
	// Product Product `json:"product_id"`
	ProductRefer uint
}
