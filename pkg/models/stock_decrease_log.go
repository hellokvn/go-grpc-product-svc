package models

import "gorm.io/gorm"

type StockDecreaseLog struct {
	gorm.Model
	Id           int32 `json:"id" gorm:"primaryKey"`
	OrderId      int32 `json:"order_id"`
	ProductRefer int32 `json:"product_id"`
}
