package gofetch

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Why I seem to need two almost identical structs:
//
// ISSUE 1: I have to use Time when unmarshaling, because time.Time
//          failes to unmarshal 2015-04-24T23:02:50 with: `parsing time
//          "2015-04-24T23:02:50" as "2006-01-02T15:04:05Z07:00": cannot parse
//          "" as "Z07:00""`
// ISSUE 2: I have to use time.Time in the gorm struct, because I can't make
//          it use Time to create & update db columns
// ISSUE 3: I have no idea how to get Rise, Set attrs without defining a Sun
//          struct, this is the XML structure
//          <forecast>
//            <sun rise="2015-04-25T06:04:04" set="2015-04-25T20:01:20"/>
//          </forecast>
// SOLUTION 1: I am now using string to parse dates in ForecastXml and then
//             just convert manually in ToForecast, this removes the need for
//             custom UnmarshalXML and UnmarshalXMLAttr functions

type Update struct {
	gorm.Model
	LastUpdate time.Time
	NextUpdate time.Time
	Rise       time.Time
	Set        time.Time
	Forecasts  []Forecast
}
type UpdateXml struct {
	LastUpdate string `xml:"meta>lastupdate"`
	NextUpdate string `xml:"meta>nextupdate"`
	Sun        struct {
		Rise string `xml:"rise,attr"`
		Set  string `xml:"set,attr"`
	} `xml:"sun"`
	Forecasts []ForecastXml `xml:"forecast>tabular>time"`
}

// Helper function for converting ForecastXml -> Forecast
func (u *UpdateXml) ToUpdate() Update {
	var forecasts []Forecast

	forecasts = make([]Forecast, len(u.Forecasts))
	for i, v := range u.Forecasts {
		forecasts[i] = v.ToForecast()
	}

	return Update{
		LastUpdate: parse(u.LastUpdate),
		NextUpdate: parse(u.NextUpdate),
		Rise:       parse(u.Sun.Rise),
		Set:        parse(u.Sun.Set),
		Forecasts:  forecasts,
	}
}
