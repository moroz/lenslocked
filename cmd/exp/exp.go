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

type Order struct {
	ID          int
	Amount      int
	Description string
	UserID      int
}

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

	// name := "Example User"
	// email := "user@example.com"
	// row := db.QueryRow(`insert into users (name, email) values ($1, $2) returning id`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created. id =", id)

	// 	id := 1
	// 	row := db.QueryRow(`select name, email from users where id = $1;`, id)
	// 	var name, email string
	// 	err = row.Scan(&name, &email)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("User information: name=%s, email=%s\n", name, email)

	// 	for i := 1; i <= 5; i++ {
	// 		amount := i * 100
	// 		desc := fmt.Sprintf("Fake order #%d", i)
	// 		_, err = db.Exec(`
	// 		insert into orders(user_id, amount, description)
	// 		values ($1, $2, $3)
	// 		`, 2, amount, desc)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}

	// 	fmt.Println("Created fake orders.")

	rows, err := db.Query(`select name, email from users where id = $1`, 2)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserID = 2
		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
	}
}
