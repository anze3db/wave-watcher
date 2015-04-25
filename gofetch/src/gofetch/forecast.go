package gofetch

import (
	"encoding/xml"
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

type Forecast struct {
	gorm.Model
	LastUpdate time.Time
	NextUpdate time.Time
	Rise       time.Time
	Set        time.Time
}
type ForecastXml struct {
	LastUpdate Time `xml:"meta>lastupdate"`
	NextUpdate Time `xml:"meta>nextupdate"`
	Sun        struct {
		Rise Time `xml:"rise,attr"`
		Set  Time `xml:"set,attr"`
	} `xml:"sun"`
}

// Helper function for converting ForecastXml -> Forecast
func (f *ForecastXml) ToForecast() Forecast {
	fo := Forecast{}
	fo.LastUpdate = f.LastUpdate.UTC()
	fo.NextUpdate = f.NextUpdate.UTC()
	fo.Rise = f.Sun.Rise.UTC()
	fo.Set = f.Sun.Set.UTC()
	return fo
}

// My Custom Time struct so that I can overwrite the default formatter
type Time struct {
	time.Time
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// We don't really care about comments or unicorns
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse("2006-01-02T15:04:05", v)
	panic(err)
	*t = Time{parse}
	return nil
}

func (t *Time) UnmarshalXMLAttr(attr xml.Attr) error {
	parse, err := time.Parse("2006-01-02T15:04:05", attr.Value)
	panic(err)
	*t = Time{parse}
	return nil
}
