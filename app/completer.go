package main

import (
	"fmt"
	"strings"
)

type Completer struct {
	commands []string
}

func (c *Completer) Do(line []rune, pos int) (newLine [][]rune, length int) {
	input := string(line[:pos])

	// only handle first word
	if strings.Contains(input, " ") {
		return nil, 0
	}

	var result [][]rune

	for _, command := range c.commands {
		if i := strings.LastIndex(command, input); i != -1 {
			end := command[len(input):]
			result = append(result, []rune(end + " "))
		}
	}

	// bell alert
	if len(result) == 0 {
		fmt.Print("\x07")
	}

	return result, len(result)
}

var completer = Completer{
	commands: []string{
		"echo",
		"exit",
	},
}
