package models

type Category struct {
	Id      int    `json:"id"`
	Libelle string `json:"libelle"`
}

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Id_category int    `json:"id_category"`
}
