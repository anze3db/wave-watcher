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

func Parse(data string) Forecast {
	f := Forecast{}
	err := xml.Unmarshal([]byte(data), &f)
	panicErr(err)
	return f
}
