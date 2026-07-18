package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var handlers map[string]func([]string)

func handleEcho(params []string) {
	fmt.Println(strings.Join(params[1:], " "))
}

func handleType(params []string) {
	if len(params) > 1 {
		checkCommand := params[1]
		if handlers[checkCommand] != nil {
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

func init() {
	handlers = map[string]func([]string){
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
				if handler := handlers[command]; handler != nil {
					handler(params)
				} else {
					fmt.Printf("%s: command not found\n", command)
				}
			}
		}
	}
}
