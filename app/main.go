package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var builtinCommands = []string{"type", "exit", "echo"}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

ShellLoop:
	for {
		fmt.Print("$ ")
		if scanner.Scan() {
			line := scanner.Text()
			params := strings.Fields(line)

			if len(params) > 0 {
				switch command := params[0]; command {
				case "exit":
					break ShellLoop
				case "echo":
					fmt.Println(strings.Join(params[1:], " "))
				case "type":
					if len(params) > 1 {
						checkCommand := params[1]
						if slices.Contains(builtinCommands, checkCommand) {
							fmt.Printf("%s is a shell builtin\n", checkCommand)
						} else {
							fmt.Printf("%s: not found\n", checkCommand)
						}
					} else {
						fmt.Println("Please input a command")
					}
				default:
					fmt.Printf("%s: command not found\n", command)
				}
			}

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
