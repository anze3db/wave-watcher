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
	panic(err)
	data, err := ioutil.ReadAll(res.Body)
	panic(err)
	return string(data)
}
