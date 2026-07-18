package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
