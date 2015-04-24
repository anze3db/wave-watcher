package gofetch

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

	s := forecast.LastUpdate
	if s != "2015-04-18T09:18:16" {
		t.Errorf("LastUpdate not parsed properly %s", s)
	}
	s = forecast.NextUpdate
	if s != "2015-04-19T22:00:00" {
		t.Errorf("NextUpdate not parsed properly %s", s)
	}
	s = forecast.Sun.Rise
	if s != "2015-04-18T06:15:37" {
		t.Errorf("SunRise not parsed properly %s", s)
	}
	s = forecast.Sun.Set
	if s != "2015-04-18T19:52:34" {
		t.Errorf("SunRise not parsed properly %s", s)
	}
}
