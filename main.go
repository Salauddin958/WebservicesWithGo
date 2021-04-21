package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pluralsight/webservices/database"
	"github.com/pluralsight/webservices/product"
	"github.com/pluralsight/webservices/receipt"
)

const basePath = "/api"

func main() {
	database.SetUpDatabase()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
