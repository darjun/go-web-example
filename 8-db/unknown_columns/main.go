package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/department")
	if err != nil {
		log.Fatal("open failed: ", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM employees")
	if err != nil {
		log.Fatal("prepare failed: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("exec failed: ", err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal("columns failed: ", err)
	}

	data := make([]interface{}, len(cols), len(cols))
	for i := range data {
		data[i] = new(string)
	}

	for rows.Next() {
		err = rows.Scan(data...)
		if err != nil {
			log.Fatal("scan failed: ", err)
		}

		for i := 0; i < len(cols); i++ {
			fmt.Printf("%s: %s ", cols[i], *(data[i].(*string)))
		}
		fmt.Println()
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}