package main

import (
	"os/exec"
)

func manOutput(arg string) (manOut []byte, e error) {
	manCmd := exec.Command("man", "-P", "cat", arg)

	manOut, e = manCmd.CombinedOutput()

	return
}
