package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var builtinCommandHandlers map[string]func([]string)

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

func init() {
	builtinCommandHandlers = map[string]func([]string){
		"echo": handleEcho,
		"type": handleType,
		"exit": handleExit,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")
		if scanner.Scan() {
			line := scanner.Text()
			params := strings.Fields(line)

			if len(params) > 0 {
				command := params[0]
				if builtinHandler := builtinCommandHandlers[command]; builtinHandler != nil {
					builtinHandler(params)
				} else {
					err := commandHandler(params)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
					}
				}
			}
		}
	}
}
