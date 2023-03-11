package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://localhost/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var greeting string
	err = db.QueryRow("select version()").Scan(&greeting)
	if err != nil {
		panic(err)
	}

	fmt.Println(greeting)
}
