package routes

import (
	"crud-api/pkg/controllers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var RegisterBookStoreRoutes = func(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/books", controllers.CreateBook(db)).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.DeleteBookById(db)).Methods("DELETE")
	router.HandleFunc("/books/{id}", controllers.UpdateBookById(db)).Methods("PUT")

}
