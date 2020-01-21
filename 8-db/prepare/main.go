package main

import (
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/department")
	if err != nil {
		log.Fatal("open failed: ", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, name, age, salary from employees where id = ?")
	if err != nil {
		log.Fatal("prepare failed: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(2)
	if err != nil {
		log.Fatal("query failed: ", err)
	}
	defer rows.Close()

	var (
		id int
		name string
		age int
		salary int
	)
	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &salary)
		if err != nil {
			log.Fatal("scan failed: ", err)
		}
		log.Printf("id:%d name:%s age:%d salary:%d\n", id, name, age, salary)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}