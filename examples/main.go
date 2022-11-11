package main

import (
	"database/sql"
	"fmt"
	_ "github.com/CN-Syndra/go-wxsqlite3"
	"log"
	"net/url"
)

func main() {

	key := url.QueryEscape("70d63bab-7022-41d8-aeda-2a9db06724af")
	dbname := fmt.Sprintf("./test.db?_db_key=%s", key)
	db, err := sql.Open("sqlite3", dbname)

	//db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	sql := `SELECT password FROM tb_account`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		name := ""
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		fmt.Println(name)
	}
}
