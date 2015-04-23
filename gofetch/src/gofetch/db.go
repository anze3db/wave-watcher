package gofetch

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func db() *sql.DB {
	dbinfo := fmt.Sprintf("postgres://postgres:%s@%s:%s?sslmode=disable",
		os.Getenv("DB_ENV_POSTGRES_PASSWORD"),
		os.Getenv("DB_PORT_5432_TCP_ADDR"),
		os.Getenv("DB_PORT_5432_TCP_PORT"))

	db, err := sql.Open("postgres", dbinfo)
	panicErr(err)
	return db
}
