package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dB, err := gorm.Open("mysql", "root:3143@/managementweb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = dB
}

func GetDB() *gorm.DB {
	return db
}
