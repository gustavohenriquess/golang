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

	db := connectDB()
	defer db.Close()
}
func Index(w http.ResponseWriter, r *http.Request) {
	products := []Product{{Name: "Mobile", Description: "Android", Price: 1000, Quantity: 10},
		{Name: "Laptop", Description: "Windows", Price: 2000, Quantity: 20},
		{Name: "TV", Description: "LED", Price: 3000, Quantity: 30},
		{"Fridge", "Refrigerator", 4000, 40},
		{"Washing Machine", "Dryer", 5000, 50},
	}
	tmpl.ExecuteTemplate(w, "Index", products)
}

func connectDB() *sql.DB {
	connection := "user=postgres dbname=alura_loja password=Postgres2018! host=teste-postgres sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}
