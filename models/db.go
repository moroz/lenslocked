package models

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(dbUrl string) (*sql.DB, error) {
	return sql.Open("pgx", dbUrl)
}
