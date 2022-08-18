package router

import (
	"product/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/products/{id}", middleware.GetProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/products", middleware.GetAllProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/products", middleware.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/products/{id}", middleware.UpdateProduct).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/products/{id}", middleware.DeleteProduct).Methods("DELETE", "OPTIONS")

	return router
}
