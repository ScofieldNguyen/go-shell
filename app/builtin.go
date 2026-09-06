package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func handleEcho(params []string, outWriter io.Writer, errWriter io.Writer) {
	fmt.Fprintln(outWriter, (strings.Join(params[1:], " ")))
}

func handleType(params []string, outWriter io.Writer, errWriter io.Writer) {
	if len(params) > 1 {
		checkCommand := params[1]
		if builtinCommandHandlers[checkCommand] != nil {
			fmt.Fprintf(outWriter, "%s is a shell builtin\n", checkCommand)
		} else {
			path, err := exec.LookPath(checkCommand)
			if err != nil {
				fmt.Fprintf(errWriter, "%s: not found\n", checkCommand)
			} else {
				fmt.Fprintf(outWriter, "%s is %s\n", checkCommand, path)
			}
		}
	} else {
		fmt.Fprintln(errWriter, "Please input a command")
	}
}

func handleExit(params []string, outWriter io.Writer, errWriter io.Writer) {
	if len(params) > 1 {
		if i, err := strconv.Atoi(params[1]); err == nil {
			os.Exit(i)
		}
	}
	os.Exit(0)
}

func handlePwd(params []string, outWriter io.Writer, errWriter io.Writer) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(errWriter, err)
		return
	}
	fmt.Fprintln(outWriter, path)
}

func handleCd(params []string, outWriter io.Writer, errWriter io.Writer) {
	var path string

	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(errWriter, err)
		return
	}

	if len(params) < 2 {
		path = homedir
	} else {
		path = params[1]
	}

	if path == "~" {
		path = homedir
	}

	err = os.Chdir(path)
	if err != nil {
		fmt.Fprintf(errWriter, "cd: %s: No such file or directory\n", path)
	}
}
