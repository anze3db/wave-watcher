package gofetch

import (
	"log"
	"time"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func panic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func print(i ...interface{}) {
	log.Printf("%s\n", i)
}

func parse(s string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05", s)
	panic(err)
	return t
}
