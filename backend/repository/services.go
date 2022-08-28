package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product/models"
	"strconv"

	"github.com/gorilla/mux"
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// insert a product in the db Mysql
func Create(w http.ResponseWriter, r *http.Request) {

	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty product of type models.Product
	var product models.Product

	// decode the json request to product
	e := json.NewDecoder(r.Body).Decode(&product)

	if e != nil {
		log.Fatalf("Unable to decode the request body.  %v", e)
	}

	// call insert product function and pass the product
	insertProduct(product)

	// send the response
	json.NewEncoder(w).Encode(product)

	// ici possiblite de changer le formet de la reponse
	//insertID := insertProduct(product)
	// format a response object
	//res := response{
	//	ID:      insertID,
	//	Message: "Product created successfully",
	//}
	// send the response
	//json.NewEncoder(w).Encode(product)

}

// GetUser will return a single user by its id
func GetProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, e := strconv.Atoi(params["id"])
	if e != nil {
		log.Fatalf("Unable to convert the string into int.  %v", e)
	}

	// call the getUser function with user id to retrieve a single user
	product, e := getProduct(int64(id))

	if e != nil {
		log.Fatalf("Unable to get user. %v", e)
	}

	// send the response
	json.NewEncoder(w).Encode(product)
}

// GetAllUser will return all the products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get all the products in the db
	products, e := getAllProducts()
	if e != nil {
		log.Fatalf("Unable to get all user. %v", e)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct update product's detail in the db
func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the product(id) from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, e := strconv.Atoi(params["id"])
	if e != nil {
		log.Fatalf("Unable to convert the string into int.  %v", e)
	}

	// create an empty product of type models.Product
	var product models.Product

	// decode the json request to product form
	e = json.NewDecoder(r.Body).Decode(&product)
	if e != nil {
		log.Fatalf("Unable to decode the request body.  %v", e)
	}

	// call update product to update the product
	updatedRows := updateProduct(int64(id), product)

	// format the message string
	msg := fmt.Sprintf("Product updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteProduct delete product's detail in the db
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the id product from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, e := strconv.Atoi(params["id"])
	if e != nil {
		log.Fatalf("Unable to convert the string into int.  %v", e)
	}

	// call the deleteProduct, convert the int to int64
	deletedRows := deleteProduct(int64(id))

	// format the message string
	msg := fmt.Sprintf("Product delete successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
