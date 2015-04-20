package fetcher

import (
	"encoding/xml"
	"log"
)

type Forecast struct {
	LastUpdate FTime `xml:"meta>lastupdate"`
	NextUpdate FTime `xml:"meta>nextupdate"`
	Sun        struct {
		Rise string `xml:"rise,attr"`
	} `xml:"sun"`
}

func ParseForecast(data string) Forecast {
	f := Forecast{}
	err := xml.Unmarshal([]byte(data), &f)
	if err != nil {
		log.Fatal("error when unmarshalling: ", err)
		return f
	}
	return f
}

func main() {
}
