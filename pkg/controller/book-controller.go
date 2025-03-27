package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Skpanchall/go-bookstore/pkg/models"
	"github.com/Skpanchall/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func CreateBook(resp http.ResponseWriter, req *http.Request) {
	book := &models.Book{}
	utils.ParseBody(req, book)
	newBooks := book.CreateBook()
	res, _ := json.Marshal(newBooks)
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)

}

func GetBook(resp http.ResponseWriter, req *http.Request) {
	books := models.GetAllBook()
	res, _ := json.Marshal(books)
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)

}
func GetBookById(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	books, _ := models.GetBookById(id)
	res, _ := json.Marshal(books)
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)

}
func UpdateBookById(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	updateBooks := models.Book{}
	utils.ParseBody(req, updateBooks)
	book, db := models.GetBookById(id)
	if updateBooks.Name != "" {
		book.Name = updateBooks.Name
	}
	if updateBooks.Author != "" {
		book.Author = updateBooks.Author
	}
	db.Save(&book)

	res, _ := json.Marshal(book)
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)

}
func DeleteBookByID(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	books := models.DeleteBookById(id)
	res, _ := json.Marshal(books)
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)

}
