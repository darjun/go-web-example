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

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("begin failed: ", err)
	}

	stmt, err := tx.Prepare("UPDATE employees SET team_id=? WHERE id=?")
	if err != nil {
		tx.Rollback()
		log.Fatal("prepare failed: ", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(2, 1)
	if err != nil {
		tx.Rollback()
		log.Fatal("exec failed: ", err)
	}

	tx.Commit()
}