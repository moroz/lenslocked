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

const migration string = `
create table if not exists users (
	id serial primary key,
	name text,
	email text not null
);

create table if not exists orders (
	id serial primary key,
	user_id int not null references users (id),
	amount int,
	description text
);
`

func main() {
	db, err := sql.Open("pgx", DBUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(migration)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")

	name := "Example User"
	email := "user@example.com"
	row := db.QueryRow(`insert into users (name, email) values ($1, $2) returning id`, name, email)
	var id int
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("User created. id =", id)
}
