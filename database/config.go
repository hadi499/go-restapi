package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-rest-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/go_api?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&models.Post{}, &models.User{})

	DB = database
}
