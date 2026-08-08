package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func commandHandler(params []string, writer io.Writer) error {
	if len(params) == 0 {
		return errors.New("no provided params")
	}

	command := params[0]

	cmd := exec.Command(command, params[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = writer

	if err := cmd.Run(); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return fmt.Errorf("%s: command not found", command)
		}
		return err
	}

	return nil
}
