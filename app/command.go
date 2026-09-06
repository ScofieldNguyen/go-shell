package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func commandHandler(params []string, outWriter io.Writer, errWriter io.Writer) {
	if len(params) == 0 {
		fmt.Fprintln(errWriter, "no provided params")
		return
	}

	command := params[0]

	cmd := exec.Command(command, params[1:]...)
	cmd.Stderr = errWriter
	cmd.Stdout = outWriter

	if err := cmd.Run(); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			fmt.Fprintf(errWriter, "%s: command not found\n", command)
		}
	}
}
