package gofetch

import "log"

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func panicErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func printErr(err error) {
	if err != nil {
		log.Print(err)
	}
}

func print(i ...interface{}) {
	log.Printf("%s\n", i)
}
