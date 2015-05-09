package gofetch

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	// Blank import needed for the db
	_ "github.com/lib/pq"
	"os"
)

var db = initDb()
var DB = db

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
	// db.session.LogMode(true)
	panic(err)
}

func (db *Db) Session() gorm.DB {
	return db.session
}
