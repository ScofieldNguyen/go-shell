package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func handleEcho(params []string, writer io.Writer) {
	fmt.Fprintln(writer, (strings.Join(params[1:], " ")))
}

func handleType(params []string, writer io.Writer) {
	if len(params) > 1 {
		checkCommand := params[1]
		if builtinCommandHandlers[checkCommand] != nil {
			fmt.Fprintf(writer, "%s is a shell builtin\n", checkCommand)
		} else {
			path, err := exec.LookPath(checkCommand)
			if err != nil {
				fmt.Fprintf(writer, "%s: not found\n", checkCommand)
			} else {
				fmt.Fprintf(writer, "%s is %s\n", checkCommand, path)
			}
		}
	} else {
		fmt.Fprintln(writer, "Please input a command")
	}
}

func handleExit(params []string, writer io.Writer) {
	if len(params) > 1 {
		if i, err := strconv.Atoi(params[1]); err == nil {
			os.Exit(i)
		}
	}
	os.Exit(0)
}

func handlePwd(params []string, writer io.Writer) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(writer, err)
	}
	fmt.Fprintln(writer, path)
}

func handleCd(params []string, writer io.Writer) {
	var path string

	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(writer, err)
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
		fmt.Fprintf(writer, "cd: %s: No such file or directory\n", path)
	}
}
