package gofetch

import (
	"encoding/xml"
)

type Forecast struct {
	LastUpdate string `xml:"meta>lastupdate"`
	NextUpdate string `xml:"meta>nextupdate"`
	Sun        struct {
		Rise string `xml:"rise,attr"`
		Set  string `xml:"set,attr"`
	} `xml:"sun"`
}

func ParseForecast(data string) Forecast {
	f := Forecast{}
	err := xml.Unmarshal([]byte(data), &f)
	panicErr(err)
	return f
}

func Run() {
	db := Db()
	defer db.Close()

	str := "INSERT INTO readings (lid, last_update, next_update, sun_rise, sun_set) values ($1, $2, $3, $4, $5);"
	stmt, err := db.Prepare(str)
	defer stmt.Close()
	panicErr(err)

	_, err = stmt.Exec(1, "2015-04-18T09:18:16", "2015-04-18T09:18:17", "2015-04-18T09:18:18", "2015-04-18T09:18:19")
	panicErr(err)

	print("Done.")
}
