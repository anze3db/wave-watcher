package gofetch

import (
	"encoding/xml"
)

func Parse(data string) Update {
	f := updateXML{}
	err := xml.Unmarshal([]byte(data), &f)
	panic(err)
	return f.ToUpdate()
}
