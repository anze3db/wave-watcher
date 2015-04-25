package gofetch

import (
	"encoding/xml"
)

func Parse(data string) Forecast {
	f := ForecastXml{}
	err := xml.Unmarshal([]byte(data), &f)
	panic(err)
	return f.ToForecast()
}
