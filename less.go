package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func colorLessEnv() [7]string {
	return [7]string{
		fmt.Sprintf("LESS_TERMCAP_mb=%s", fmtAnsi(bold, blue)),
		fmt.Sprintf("LESS_TERMCAP_so=%s", fmtAnsi(bold, green)),
		fmt.Sprintf("LESS_TERMCAP_md=%s", fmtAnsi(red)),
		fmt.Sprintf("LESS_TERMCAP_us=%s", fmtAnsi(underline, yellow)),
		fmt.Sprintf("LESS_TERMCAP_me=%s", fmtAnsi(reset)),
		fmt.Sprintf("LESS_TERMCAP_se=%s", fmtAnsi(reset)),
		fmt.Sprintf("LESS_TERMCAP_ue=%s", fmtAnsi(reset)),
	}
}

func lessCommand(userEnv []string, out, err io.WriteCloser) *exec.Cmd {
	lessEnv := colorLessEnv()

	env := append(lessEnv[:], userEnv[:]...)

	lessCmd := exec.Command("less")
	lessCmd.Stdout = os.Stdout
	lessCmd.Stderr = os.Stderr

	lessCmd.Env = env

	return lessCmd
}
