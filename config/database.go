package config

import (
	"github.com/mhdianrush/go-products-rest-api/entities"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	logger := logrus.New()

	db, err := gorm.Open(mysql.Open("root:admin@tcp(127.0.0.1:3306)/go_products_rest_api?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	logger.Println("Database Connected")
	db.AutoMigrate(&entities.Product{})

	DB = db
}
