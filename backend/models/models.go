package models

type Category struct {
	Id      int    `json:"id"`
	Libelle string `json:"libelle"`
}

type Product struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price,string"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity,string"`
	Id_category int    `json:"id_category,string"`
}
