package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func openDB() {
	db, err := sql.Open("mysql", `orderservice:1234@order_service`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
