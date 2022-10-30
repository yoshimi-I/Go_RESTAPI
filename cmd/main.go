package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Db *sql.DB

func main() {
	Db, err := sql.Open("sqlite3", "db/example.sql")
	if err != nil {
		log.Fatal("残念")
	}
	defer Db.Close()

	cmd := `CREATE TABLE IF NOT EXISTS persons(name STRING,age INT)`

	_, err = Db.Exec(cmd)
	if err != nil {
		log.Fatalln("ざまあ")
	}

}
