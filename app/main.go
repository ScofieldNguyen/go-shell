package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")
		if scanner.Scan() {
			line := scanner.Text()
			params := strings.Fields(line)

			if len(params) > 0 {
				command := params[0]

				if command == "exit" {
					break
				}

				fmt.Printf("%s: command not found\n", command)
			}

			if err := scanner.Err(); err != nil {
				break
			}
		}
	}
}
