package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	leader := color("> ", bold, green)

	prompt := fmt.Sprintf("Input a command to flannotate:\n%s", leader)
	fmt.Print(color(prompt, bold, blue))

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	cmd := strings.TrimSpace(input)

	if err != nil {
		return err
	}

	prompt = fmt.Sprintf("Input your flannotation for %s:\n%s",
		color(cmd, bold, red), leader)

	fmt.Print(color(prompt, bold, blue))

	input, err = reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Println(color("🍮", magenta))
	return nil
}
