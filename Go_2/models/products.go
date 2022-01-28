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

	selectAllProducs, err := db.Query("SELECT * FROM products order by id asc")

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

func GetProduct(id string) Product {
	db := db.ConnectDB()

	selectProduct, err := db.Query("SELECT * FROM products WHERE id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for selectProduct.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProduct.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = nome
		product.Description = descricao
		product.Price = preco
		product.Quantity = quantidade

	}

	defer db.Close()
	return product
}

func UpdateProduct(id int, name, description string, quantity int, price float64) {
	db := db.ConnectDB()

	updateProduct, err := db.Prepare("UPDATE products SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
