package fetcher

import (
	"testing"
)

func TestParseForecast(t *testing.T) {
	data := `
	<?xml version="1.0" encoding="utf-8"?>
	<weatherdata>
	  <meta>
	    <lastupdate>2015-04-18T09:18:16</lastupdate>
	    <nextupdate>2015-04-19T22:00:00</nextupdate>
	  </meta>
	  <sun rise="2015-04-18T06:15:37" set="2015-04-18T19:52:34" />
	  <forecast>
	  </forecast>
	</weatherdata>
	`

	forecast := ParseForecast(data)

	d := forecast.LastUpdate
	if d.Day() != 18 && d.Month() != 4 && d.Year() != 2015 {
		t.Error("LastUpdate not parsed properly")
	}
	d = forecast.NextUpdate
	if d.Day() != 19 && d.Month() != 4 && d.Year() != 2015 {
		t.Error("NextUpdate not parsed properly")
	}
	s := forecast.Sun.Rise
	if s != "2015-04-18T06:15:37" {
		t.Errorf("SunRise not parsed properly %s", d)
	}
}
