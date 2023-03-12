package models

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {
	return sql.Open("pgx", os.Getenv("DATABASE_URL"))
}
