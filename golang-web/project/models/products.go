package models

import (
	"database/sql"

	"github.com/userbarbosa/golang-alura/golang-web/project/v2/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func ListProducts() []Product {
	db := db.ConnectDatabase()
	defer db.Close()

	productsQuery, err := db.Query("select * from products order by id")
	if err != nil {
		panic(err.Error())
	}

	products := []Product{}

	for productsQuery.Next() {
		products = append(products, scanProduct(productsQuery))
	}

	return products
}

func NewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()
	defer db.Close()

	insertCommand, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertCommand.Exec(name, description, price, quantity)
}

func DeleteProductFromId(id string) {
	db := db.ConnectDatabase()
	defer db.Close()

	deleteCommand, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteCommand.Exec(id)
}

func GetProduct(id string) Product {
	db := db.ConnectDatabase()
	defer db.Close()

	query, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productInstance := Product{}
	for query.Next() {
		productInstance = scanProduct(query)
	}
	return productInstance
}

func scanProduct(query *sql.Rows) (productInstance Product) {
	var id, quantity int
	var name, description string
	var price float64

	err := query.Scan(&id, &name, &description, &price, &quantity)
	if err != nil {
		panic(err.Error())
	}

	productInstance.Id = id
	productInstance.Name = name
	productInstance.Description = description
	productInstance.Price = price
	productInstance.Quantity = quantity

	return productInstance
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()
	defer db.Close()

	updateCommand, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateCommand.Exec(name, description, price, quantity, id)
}
