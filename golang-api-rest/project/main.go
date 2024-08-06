package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/userbarbosa/golang-alura/golang-api-rest/project/v2/database"
	"github.com/userbarbosa/golang-alura/golang-api-rest/project/v2/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	database.ConnectDatabase()
	apiPort := os.Getenv("API_PORT")
	fmt.Println("Initialing GO Server at ", apiPort)
	routes.HandleRequest()
}
