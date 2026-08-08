package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func commandHandler(params []string, writer io.Writer) {
	if len(params) == 0 {
		fmt.Fprintln(os.Stderr, "no provided params")
		return
	}

	command := params[0]

	cmd := exec.Command(command, params[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = writer

	if err := cmd.Run(); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
		}
	}
}
