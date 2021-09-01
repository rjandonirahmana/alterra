package config

import (
	model "mvcApi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func NewDb() *gorm.DB {

	var err error
	dsn := ":@tcp(127.0.0.1:3306)/alterra?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&model.Users{})
	Db.AutoMigrate(&model.Book{})

	return Db

}
