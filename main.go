package main

import (
	"log"
	"os"
)

func main() {
	var err error

	args := os.Args[1:]

	switch len(args) {
	case 0:
		err = flannotate()
	case 1:
		err = flanpage(args[0])
	default:
		log.Fatal("Pass 1 or 0 args")
	}

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
