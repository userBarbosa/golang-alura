package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/userbarbosa/golang-alura/golang-gin/project/v2/models"
	"github.com/userbarbosa/golang-alura/golang-gin/project/v2/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	models.Students = []models.Student{
		{Name: "Martin Noah", IdentificationNumber: "74937565540", RegistrationNumber: "214383532"},
		{Name: "Mariane Silva", IdentificationNumber: "26856909802", RegistrationNumber: "501936567"},
		{Name: "Rodrigo Gustavo", IdentificationNumber: "10880031840", RegistrationNumber: "251980650"},
	}

	apiPort := os.Getenv("API_PORT")
	fmt.Println("Initialing GO Server at ", apiPort)
	routes.HandleRequests()
}
