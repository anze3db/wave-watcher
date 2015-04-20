package fetcher

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Db() *sql.DB {
	dbinfo := fmt.Sprintf("postgres://postgres:%s@%s:%s?sslmode=disable",
		os.Getenv("DB_ENV_POSTGRES_PASSWORD"),
		os.Getenv("DB_PORT_5432_TCP_ADDR"),
		os.Getenv("DB_PORT_5432_TCP_PORT"))

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
