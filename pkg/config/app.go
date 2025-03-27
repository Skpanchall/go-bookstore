package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "username:password@tcp(localhost:3306)/go_bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	fmt.Println("open successfully ")
	db = d

}

func GetDB() *gorm.DB {
	return db
}
