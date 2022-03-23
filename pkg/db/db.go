package db

import (
	"log"

	"github.com/hellokvn/go-grpc-product-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://kevin@localhost:5432/tasks"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.StockDecreaseLog{})

	return db
}
