package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {

	db := connectDB()

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
	tmpl.ExecuteTemplate(w, "Index", products)

	defer db.Close()
}

func connectDB() *sql.DB {
	connection := "user=postgres dbname=alura_loja password=Postgres2018! host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}
