package main

import (
	"log"
	"net/http"
	"product/router"
)

func main() {

	r := router.Router()

	log.Println("Server started on: http://localhost:4200")

	log.Fatal(http.ListenAndServe(":4200", r))

}
