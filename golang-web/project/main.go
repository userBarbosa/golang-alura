package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/userbarbosa/golang-alura/golang-web/project/v2/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
