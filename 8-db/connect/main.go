package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/not_exist_department")
	if err != nil {
		log.Fatal("connect database failed: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("ping failed: ", err)
	}
}