package models

import (
	"golang/Go_2/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SearchAllProducts() []Product {
	db := db.ConnectDB()

	selectAllProducs, err := db.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducs.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducs.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = nome
		p.Description = descricao
		p.Price = preco
		p.Quantity = quantidade
		products = append(products, p)
	}

	defer db.Close()

	return products
}

func InsertProduct(name, description string, quantityConverted int, priceConverted float64) {

	db := db.ConnectDB()

	insertData, err := db.Prepare("INSERT INTO products(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, priceConverted, quantityConverted)

	defer db.Close()

}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

	defer db.Close()
}
