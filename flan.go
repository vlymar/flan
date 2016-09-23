package main

import (
	"fmt"
	//"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
$ lsof -i :6419
$ flan "list processes sitting on $port$6419$"

$ flan lsof
# list processes sitting on $port
lsof -i :$port
*/

// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
// https://github.com/fatih/color

const esc = "\x1b"

type AnsiCode int

const (
	reset AnsiCode = iota
	bold
	faint
	italic
	underline
)

const (
	black AnsiCode = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

func fmtAnsi(codes ...AnsiCode) string {
	codeStrings := make([]string, len(codes))
	for i, v := range codes {
		codeStrings[i] = strconv.Itoa(int(v))
	}
	delimitedCodes := strings.Join(codeStrings, ";")
	return fmt.Sprintf("%s[%sm", esc, delimitedCodes)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("no args passed")
	}

	colorLessEnv := [7]string{
		fmt.Sprintf("LESS_TERMCAP_mb=%s", fmtAnsi(bold, blue)),
		fmt.Sprintf("LESS_TERMCAP_so=%s", fmtAnsi(bold, green)),
		fmt.Sprintf("LESS_TERMCAP_md=%s", fmtAnsi(red)),
		fmt.Sprintf("LESS_TERMCAP_us=%s", fmtAnsi(underline, yellow)),
		fmt.Sprintf("LESS_TERMCAP_me=%s", fmtAnsi(reset)),
		fmt.Sprintf("LESS_TERMCAP_se=%s", fmtAnsi(reset)),
		fmt.Sprintf("LESS_TERMCAP_ue=%s", fmtAnsi(reset)),
	}
	userEnv := os.Environ()
	env := append(colorLessEnv[:], userEnv[:]...)

	manCmd := exec.Command("man", "-P", "cat", args[0])

	manCmd.Stderr = os.Stderr
	manOut, err := manCmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	lessCmd := exec.Command("less")
	lessIn, _ := lessCmd.StdinPipe()
	lessCmd.Env = env
	lessCmd.Stdout = os.Stdout
	lessCmd.Stderr = os.Stderr

	if err := lessCmd.Start(); err != nil {
		log.Fatal(err)
	}

	lessIn.Write([]byte("future flannotation...\n\n"))
	lessIn.Write(manOut)
	lessIn.Close()

	if err := lessCmd.Wait(); err != nil {
		log.Fatal(err)
	} else {
		os.Exit(0)
	}
}
