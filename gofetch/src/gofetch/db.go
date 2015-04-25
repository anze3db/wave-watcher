package gofetch

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

var db = initDb()

func initDb() Db {
	d := Db{}
	d.Init()
	return d
}

type Db struct {
	session gorm.DB
}

func (db *Db) Init() {

	dbinfo := fmt.Sprintf("postgres://postgres:%s@%s:%s?sslmode=disable",
		os.Getenv("DB_ENV_POSTGRES_PASSWORD"),
		os.Getenv("DB_PORT_5432_TCP_ADDR"),
		os.Getenv("DB_PORT_5432_TCP_PORT"))

	conn, err := sql.Open("postgres", dbinfo)
	panic(err)

	db.session, err = gorm.Open("postgres", conn)
	panic(err)
}
