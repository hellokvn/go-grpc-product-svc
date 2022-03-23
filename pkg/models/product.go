package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id                int32            `json:"id" gorm:"primaryKey"`
	Name              string           `json:"name"`
	Stock             int32            `json:"stock"`
	Price             int32            `json:"price"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
