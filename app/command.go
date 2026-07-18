package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func commandHandler(params []string) error {
	if len(params) == 0 {
		return errors.New("no provided params")
	}

	command := params[0]

	cmd := exec.Command(command, params[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return fmt.Errorf("%s: command not found", command)
		}
		return err
	}

	return nil
}
