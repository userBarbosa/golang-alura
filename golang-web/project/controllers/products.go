package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/userbarbosa/golang-alura/golang-web/project/v2/models"
)

var templateInstance = template.Must(template.ParseGlob("templates/*.html"))

func ListProducts(w http.ResponseWriter, r *http.Request) {
	listProducts := models.ListProducts()
	templateInstance.ExecuteTemplate(w, "Products", listProducts)
}

func NewProductForm(w http.ResponseWriter, r *http.Request) {
	templateInstance.ExecuteTemplate(w, "New_product", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")

	priceAsFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Error at float conversion", err)
	}

	quantityAsInt, err := strconv.Atoi(quantity)
	if err != nil {
		log.Println("Error at integer conversion", err)
	}

	models.NewProduct(name, description, priceAsFloat, quantityAsInt)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("id") {
		log.Println("unable to delete", r.URL)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	id := r.URL.Query().Get("id")
	models.DeleteProductFromId(id)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("id") {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	id := r.URL.Query().Get("id")
	product := models.GetProduct(id)
	templateInstance.ExecuteTemplate(w, "Edit_product", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")

	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error at integer (id) conversion", err)
	}
	quantityAsInt, err := strconv.Atoi(quantity)
	if err != nil {
		log.Println("Error at integer (quantity) conversion", err)
	}
	priceAsFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Error at float (price) conversion", err)
	}

	models.UpdateProduct(idAsInt, name, description, priceAsFloat, quantityAsInt)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
