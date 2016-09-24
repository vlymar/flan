package main

import (
	"fmt"
	"log"
	"os"
)

/*
$ lsof -i :6419
$ flan "list processes sitting on $port$6419$"

$ flan lsof
# list processes sitting on $port
lsof -i :$port
*/

func flanpage(arg string) error {
	manOut, err := manOutput(arg, os.Stderr)
	if err != nil {
		return err
	}

	lessCmd := lessCommand(os.Environ(), os.Stdout, os.Stderr)

	lessIn, err := lessCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := lessCmd.Start(); err != nil {
		return err
	}

	lessIn.Write([]byte("future flannotation...\n\n"))
	lessIn.Write(manOut)
	lessIn.Close()

	if err := lessCmd.Wait(); err != nil {
		return err
	}

	return nil
}

func flannotate() error {
	fmt.Println("flannotating lol")
	return nil
}

func main() {
	var err error = nil

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
