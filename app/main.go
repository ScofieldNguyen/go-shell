package main

import (
	"fmt"
	"os"
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
	for {
		fmt.Print("$ ")
		if params, err := getParams(); err == nil {
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
