package models

import (
	"golang/Go_2/db"
)

type Product struct {
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

		p.Name = nome
		p.Description = descricao
		p.Price = preco
		p.Quantity = quantidade
		products = append(products, p)
	}

	defer db.Close()

	return products
}
