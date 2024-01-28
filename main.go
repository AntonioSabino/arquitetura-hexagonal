package main

import (
	"database/sql"

	db2 "github.com/AntonioSabino/go-hexagonal/adapters/db"
	"github.com/AntonioSabino/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	defer db.Close()

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product 1", 10.00)

	productService.Enable(product)
}
