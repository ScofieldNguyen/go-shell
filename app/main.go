package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var builtinCommandHandlers map[string]func([]string)

func init() {
	builtinCommandHandlers = map[string]func([]string){
		"echo": handleEcho,
		"type": handleType,
		"exit": handleExit,
		"pwd":  handlePwd,
		"cd":   handleCd,
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
