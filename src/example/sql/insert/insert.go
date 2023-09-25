package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:root@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values (?.?)")

	stmt.Exec(4000, "Bia")
	//Esperando erro
	_, erro := stmt.Exec(1, "Tiago")
	if erro != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}
