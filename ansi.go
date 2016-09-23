package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Useful Reading:
	https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
	https://github.com/fatih/color
*/

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
