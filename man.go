package main

import (
	"io"
	"os/exec"
)

func manOutput(arg string, err io.WriteCloser) (manOut []byte, e error) {
	manCmd := exec.Command("man", "-P", "cat", arg)

	manCmd.Stderr = err
	manOut, e = manCmd.Output()

	return
}
