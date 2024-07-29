package routes

import (
	"net/http"

	"github.com/userbarbosa/golang-alura/golang-web/project/v2/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.ListProducts)
	http.HandleFunc("/new", controllers.NewProductForm)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/edit", controllers.GetProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
}
