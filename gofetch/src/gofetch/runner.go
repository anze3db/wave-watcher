package gofetch

import (
	"database/sql"
	"time"
)

func Run() {
	next_update := FindLatest(Parse(Fetch()))
	duration := next_update.Sub(time.Now())
	if duration.Minutes() < 5 {
		duration = 5 * time.Minute
	}
	print("Next run in", duration)
	<-time.After(duration)
	Run()
}

func insert(f Forecast) {
	db := db()
	defer db.Close()

	str := "INSERT INTO readings (lid, last_update, next_update, sun_rise, sun_set) values ($1, $2, $3, $4, $5);"
	stmt, err := db.Prepare(str)
	defer stmt.Close()
	panicErr(err)

	_, err = stmt.Exec(1, f.LastUpdate, f.NextUpdate, f.Sun.Rise, f.Sun.Set)
	panicErr(err)

	print("Updated")
}

func FindLatest(f Forecast) time.Time {
	db := db()
	defer db.Close()

	var next_update time.Time
	err := db.QueryRow("SELECT next_update FROM readings WHERE last_update >= $1;", f.LastUpdate).Scan(&next_update)
	if err == sql.ErrNoRows {
		insert(f)
		next_update, _ = time.Parse("2006-01-02T15:04:05", f.NextUpdate)
	} else {
		panicErr(err)
	}
	return next_update
}
