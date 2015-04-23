package gofetch

import (
	"io/ioutil"
	"net/http"
)

const (
	URL = "http://www.yr.no/place/Croatia/Istria/Medulin/forecast.xml"
)

func Fetch() string {
	res, err := http.Get(URL)
	panicErr(err)
	data, err := ioutil.ReadAll(res.Body)
	panicErr(err)
	return string(data)
}
