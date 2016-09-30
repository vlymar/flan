package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func colorLessEnv() [7]string {
	return [7]string{
		fmt.Sprintf("LESS_TERMCAP_mb=%s", ansi(bold, blue)),
		fmt.Sprintf("LESS_TERMCAP_so=%s", ansi(bold, green)),
		fmt.Sprintf("LESS_TERMCAP_md=%s", ansi(red)),
		fmt.Sprintf("LESS_TERMCAP_us=%s", ansi(underline, yellow)),
		fmt.Sprintf("LESS_TERMCAP_me=%s", ansi(reset)),
		fmt.Sprintf("LESS_TERMCAP_se=%s", ansi(reset)),
		fmt.Sprintf("LESS_TERMCAP_ue=%s", ansi(reset)),
	}
}

func lessCommand(userEnv []string, out, err io.WriteCloser) *exec.Cmd {
	lessEnv := colorLessEnv()

	env := append(lessEnv[:], userEnv[:]...)

	lessCmd := exec.Command("less", "-R")
	lessCmd.Stdout = os.Stdout
	lessCmd.Stderr = os.Stderr

	lessCmd.Env = env

	return lessCmd
}
