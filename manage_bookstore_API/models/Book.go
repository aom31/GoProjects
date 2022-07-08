package models

import (
	"github.com/example/manageapi/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name        string ` gorm: "" json:"name" `
	Auther      string ` json:"auther" `
	Publication string `  json:"publication" `
}

func init() {
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Book{})

}
