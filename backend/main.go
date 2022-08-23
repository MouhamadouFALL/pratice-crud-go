package main

import (
	"log"
	"net/http"
	"product/router"
)

func main() {

	r := router.Router()

	log.Println("Server started on: http://localhost:9595")

	//handler := cors.Default().Handler(r)

	//log.Fatal(http.ListenAndServe(":9595", handler))

	log.Fatal(http.ListenAndServe(":9595", r))

}
