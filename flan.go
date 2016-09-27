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

// TODO: ctrl-c'ing out of here breaks prompt for a sec
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

	commands, err := ReadFlanFile()
	if err != nil {
		return err
	}
	flannotations, prs := commands[arg]
	if prs {
		for _, flanno := range flannotations {
			// TODO: better formatting, fix color escaping
			lessIn.Write([]byte("$ "))
			lessIn.Write([]byte(flanno[0]))
			lessIn.Write([]byte("\n"))
			lessIn.Write([]byte("# "))
			lessIn.Write([]byte(flanno[1]))
			lessIn.Write([]byte("\n\n"))
		}
	}

	lessIn.Write(manOut)
	lessIn.Close()

	if err := lessCmd.Wait(); err != nil {
		return err
	}

	return nil
}

func flannotate() error {
	leader := color("> ", bold, green)

	commands, err := ReadFlanFile()
	if err != nil {
		return err
	}

	prompt := fmt.Sprintf("Input a command example to flannotate:\n%s", leader)
	fmt.Print(color(prompt, bold, blue))

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return err
	}

	cmdEx := strings.TrimSpace(input)
	cmdName := strings.Split(cmdEx, " ")[0]

	prompt = fmt.Sprintf("Input a new flannotation for %s:\n%s",
		color(cmdEx, bold, red), leader)

	fmt.Print(color(prompt, bold, blue))

	input, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	cmdAnno := strings.TrimSpace(input)

	Store(cmdName, cmdEx, cmdAnno, commands)
	if err = WriteFlanFile(commands); err != nil {
		return err
	}

	fmt.Println(color("🍮", magenta))
	return nil
}
