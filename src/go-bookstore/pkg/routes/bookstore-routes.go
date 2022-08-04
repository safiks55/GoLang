package routes

import (
	"github.com/gorilla/mux"
	"github.com/safik/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")          // trying to create new book
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")              // get all books
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")  // get books by ID
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")   // update book using ID
	router.HandleFunc("book/{bookId}", controllers.DeleteBook).Methods("DELETE") // del particular book using id
}
