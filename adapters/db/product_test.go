package db_test

import (
	"database/sql"
	"testing"

	"github.com/AntonioSabino/go-hexagonal/adapters/db"
	"github.com/AntonioSabino/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `
		CREATE TABLE products (
			"id" string,
			"name" string,
			"status" string,
			"price" float
		);
	`

	stmt, err := Db.Prepare(table)
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `
		INSERT INTO products VALUES (
			"1",
			"Product 1",
			"disabled",
			10.00
		);
	`

	stmt, err := Db.Prepare(insert)
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("1")

	require.Nil(t, err)
	require.Equal(t, "1", product.GetID())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
	require.Equal(t, 10.00, product.GetPrice())
}

func TestProductDb_Save(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product 2"
	product.Price = 20.00

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, "disabled", productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}
