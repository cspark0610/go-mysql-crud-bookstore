package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectToDatabase() {
	dbResult, err := gorm.Open("mysql", "root:root@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	db = dbResult

}

func GetDB() *gorm.DB {
	return db
}
