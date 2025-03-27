package routes

import (
	"github.com/Skpanchall/go-bookstore/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.UpdateBookById).Methods("POST")
	router.HandleFunc("/book/{bookId}", controller.DeleteBookByID).Methods("DELETE")

}
