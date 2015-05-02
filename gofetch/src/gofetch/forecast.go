package gofetch

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Forecast is a struct that holds weather data for a given time period
type Forecast struct {
	gorm.Model
	UpdateID          int `sql:"index"`
	From              time.Time
	To                time.Time
	WindDirection     float32
	WindDirectionName string
	WindSpeed         float32
	WindSpeedName     string
	Temperature       float32
	TemperatureName   string
	Pressure          float32
	PressureName      string
	Symbol            float32
	SymbolName        string
}

// TODO: The code below is just boilerplate because I don't know how
//       to access ie <winddirection deg=""> deg without a winddirection
//       struct. I also don't know how to correctly parse time.Time
type forecastXML struct {
	From          string `xml:"from,attr"`
	To            string `xml:"to,attr"`
	WindDirection struct {
		Deg  float32 `xml:"deg,attr"`
		Code string  `xml:"code,attr"`
	} `xml:"windDirection"`
	WindSpeed struct {
		Mps  float32 `xml:"mps,attr"`
		Name string  `xml:"name,attr"`
	} `xml:"windSpeed"`
	Temperature struct {
		Value float32 `xml:"value,attr"`
		Name  string  `xml:"unit,attr"`
	} `xml:"temperature"`
	Pressure struct {
		Value float32 `xml:"value,attr"`
		Name  string  `xml:"unit,attr"`
	} `xml:"pressure"`
	Symbol struct {
		Number float32 `xml:"number,attr"`
		Name   string  `xml:"name,attr"`
	} `xml:"symbol"`
}

func (f *forecastXML) ToForecast() Forecast {
	return Forecast{
		From:              parse(f.From),
		To:                parse(f.To),
		WindDirection:     f.WindDirection.Deg,
		WindDirectionName: f.WindDirection.Code,
		WindSpeed:         f.WindSpeed.Mps,
		WindSpeedName:     f.WindSpeed.Name,
		Temperature:       f.Temperature.Value,
		TemperatureName:   f.Temperature.Name,
		Pressure:          f.Pressure.Value,
		PressureName:      f.Pressure.Name,
		Symbol:            f.Symbol.Number,
		SymbolName:        f.Symbol.Name,
	}
}
