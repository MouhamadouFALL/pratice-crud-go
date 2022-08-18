package router

import (
	"product/repository"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/products/{id}", repository.GetProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/products", repository.GetAllProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/products", repository.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/products/{id}", repository.UpdateProduct).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/products/{id}", repository.DeleteProduct).Methods("DELETE", "OPTIONS")

	return router
}
