package main

import (
	"html/template"
	"net/http"

	"golang/Go_2/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.SearchAllProducts()
	tmpl.ExecuteTemplate(w, "Index", allProducts)
}
