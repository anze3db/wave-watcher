package gofetch

import "log"

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
