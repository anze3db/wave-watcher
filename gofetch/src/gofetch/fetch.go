package gofetch

import (
	"io/ioutil"
	"net/http"
)

const (
	url = "http://www.yr.no/place/Croatia/Istria/Medulin/forecast.xml"
)

func Fetch() string {
	res, err := http.Get(url)
	panic(err)
	data, err := ioutil.ReadAll(res.Body)
	panic(err)
	return string(data)
}
