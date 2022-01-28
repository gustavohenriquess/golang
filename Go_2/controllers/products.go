package controllers

import (
	"golang/Go_2/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.SearchAllProducts()
	tmpl.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		quantity := r.FormValue("quantidade")
		price := r.FormValue("preco")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price", err)
		}
		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity", err)
		}

		models.InsertProduct(name, description, quantityConverted, priceConverted)

		http.Redirect(w, r, "/", 301)
	}

	tmpl.ExecuteTemplate(w, "New", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteProduct(id)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	product := models.GetProduct(id)
	tmpl.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		quantity := r.FormValue("quantidade")
		price := r.FormValue("preco")

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error converting id", err)
		}
		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price", err)
		}
		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting quantity", err)
		}

		models.UpdateProduct(idConverted, name, description, quantityConverted, priceConverted)

	}
	http.Redirect(w, r, "/", 301)
}
