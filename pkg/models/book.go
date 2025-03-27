package models

import (
	"github.com/Skpanchall/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var (
	db *gorm.DB
)

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book {
	db.NewRecord(&b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	book := Book{}
	db := db.Where("id = ?", id).Find(&book)
	return &book, db
}

func DeleteBookById(id int64) Book {
	book := Book{}
	db.Where("id = ?", id).Delete(book)
	return book
}
