package database

import (
	"database/sql"
	"log"
	"time"
)

var DbConn *sql.DB

func SetUpDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
