package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DBUrl string

func init() {
	DBUrl = os.Getenv("DATABASE_URL")
	if DBUrl == "" {
		panic("Environment variable DATABASE_URL is not set!")
	}
}

func main() {
	db, err := sql.Open("pgx", DBUrl)
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
