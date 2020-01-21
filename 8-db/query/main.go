package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/department")
	if err != nil {
		log.Fatal("open database failed: ", err)
	}
	defer db.Close()

	var id int
	var name string
	var age int
	var salary int
	var teamId int

	rows, err := db.Query("select id, name, age, salary, team_id from employees where id = ?", 1)
	if err != nil {
		log.Fatal("query failed: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &salary, &teamId)
		if err != nil {
			log.Fatal("scan failed: ", err)
		}
		log.Printf("id: %d name:%s age:%d salary:%d teamId:%d\n", id, name, age, salary, teamId)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}