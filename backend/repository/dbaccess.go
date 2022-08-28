package repository

import (
	"database/sql"
	"fmt"
	"log"
	"product/models"

	_ "github.com/go-sql-driver/mysql"
)

// create connection with Mysql db
func dbConnection() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "gocrud"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected!")

	// return the connection
	return db
}

// insert a product in the DB(gocrud)
func insertProduct(product models.Product) int64 {

	// create the db connection
	db := dbConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	query := "Insert Into product (name, price, description, quantity, id_category) Values (?, ?, ?, ?, ?)"

	// prepare the query
	stmt, e := db.Prepare(query)
	if e != nil {
		panic(e.Error())
	}

	// execute
	res, e := stmt.Exec(product.Name, product.Price, product.Description, product.Quantity, product.Id_category)
	if e != nil {
		panic(e.Error())
	}

	// get the last id product inserted and will return it
	id, e := res.LastInsertId()
	if e != nil {
		log.Fatalf("Unable to execute the query. %v", e)
	}

	// optional statement
	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// get one user from the DB by its userid
func getProduct(id int64) (models.Product, error) {

	// get connection
	db := dbConnection()

	// close the db connection
	defer db.Close()

	var product models.Product

	// create the select sql query
	query := "SELECT * FROM product WHERE id=?"

	// execute the sql statement
	row := db.QueryRow(query, id)

	// unmarshal the row object to user
	e := row.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Quantity, &product.Id_category)

	// optional statement
	//fmt.Printf("Product Id : ", product.Id)

	switch e {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("Unable to scan the row. %v", e)
	}

	// return empty product on error
	return product, e
}

// get one product from the DB by its id (product id)
func getAllProducts() ([]models.Product, error) {

	// create the connection
	db := dbConnection()

	// close the db connection
	defer db.Close()

	var products []models.Product

	// create the select sql query
	query := "SELECT * FROM product"

	// execute the sql statement
	rows, e := db.Query(query)
	if e != nil {
		log.Fatalf("Unable to execute the query. %v", e)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var product models.Product

		// unmarshal the row object to user
		e = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Quantity, &product.Id_category)

		if e != nil {
			log.Fatalf("Unable to scan the row. %v", e)
		}

		// append the user in the users slice
		products = append(products, product)

	}

	// return empty user on error
	return products, e
}

// update product in the DB
func updateProduct(id int64, product models.Product) int64 {

	// create the Mysql db connection
	db := dbConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	query := "Update product Set name=?, price=?, description=?, quantity=?, id_category=? WHERE id=?"

	// prepare the query
	stmt, e := db.Prepare(query)
	if e != nil {
		panic(e.Error())
	}

	// execute
	res, e := stmt.Exec(product.Name, product.Price, product.Description, product.Quantity, product.Id_category, id)
	if e != nil {
		panic(e.Error())
	}

	rowsAffected, e := res.RowsAffected()
	if e != nil {
		log.Fatalf("Error while checking the affected rows. %v", e)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete product in the DB
func deleteProduct(id int64) int64 {

	// create db connection
	db := dbConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	query := "Delete FROM product Where id=?"

	// execute the sql statement
	res, e := db.Exec(query, id)
	if e != nil {
		log.Fatalf("Unable to execute the query. %v", e)
	}

	// check how many rows affected
	rowsAffected, e := res.RowsAffected()
	if e != nil {
		log.Fatalf("Error while checking the affected rows. %v", e)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
