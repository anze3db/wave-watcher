package fetcher

import (
	"encoding/xml"
	"log"
	"time"
)

// Custom struct needed for Unmarshalling
type FTime struct {
	time.Time
}

func (c *FTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "2006-01-02T15:04:05"
	var v string
	d.DecodeElement(&v, &start)
	log.Print("PARSING", v)
	parse, err := time.Parse(format, v)
	log.Print("IN FTIME: ", parse, err, format, v)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	*c = FTime{parse}
	return nil
}
