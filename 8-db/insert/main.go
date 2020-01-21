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

	stmt, err := db.Prepare("INSERT INTO employees(name, age, salary, team_id) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal("prepare failed: ", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("柳十四", 32, 5000, 1)
	if err != nil {
		log.Fatal("exec failed: ", err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal("fetch last insert id failed: ", err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal("fetch rows affected failed: ", err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}