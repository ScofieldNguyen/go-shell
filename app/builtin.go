package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func handleEcho(params []string) {
	fmt.Println(strings.Join(params[1:], " "))
}

func handleType(params []string) {
	if len(params) > 1 {
		checkCommand := params[1]
		if builtinCommandHandlers[checkCommand] != nil {
			fmt.Printf("%s is a shell builtin\n", checkCommand)
		} else {
			path, err := exec.LookPath(checkCommand)
			if err != nil {
				fmt.Printf("%s: not found\n", checkCommand)
			} else {
				fmt.Printf("%s is %s\n", checkCommand, path)
			}
		}
	} else {
		fmt.Println("Please input a command")
	}
}

func handleExit(params []string) {
	if len(params) > 1 {
		if i, err := strconv.Atoi(params[1]); err == nil {
			os.Exit(i)
		}
	}
	os.Exit(0)
}

func handlePwd(params []string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
}

func handleCd(params []string) {
	var path string

	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
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
		fmt.Printf("cd: %s: No such file or directory\n", path)
	}
}
