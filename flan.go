package main

import (
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

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("no args passed")
	}

	manOut, err := manOutput(args[0], os.Stderr)
	if err != nil {
		log.Fatal(err)
	}

	lessCmd := lessCommand(os.Environ(), os.Stdout, os.Stderr)
	if err := lessCmd.Start(); err != nil {
		log.Fatal(err)
	}

	lessIn, err := lessCmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	lessIn.Write([]byte("future flannotation...\n\n"))
	lessIn.Write(manOut)

	if err := lessCmd.Wait(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
