package main

import (
	"fmt"
	"os"
	"strings"
)

type Completer struct {
	commands []string
}

func (c *Completer) Do(line []rune, pos int) (newLine [][]rune, length int) {
	input := string(line[:pos])
	input = strings.Trim(input, " ")

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

func listExecutables(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var executables []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		if info.Mode()&0111 != 0 {
			executables = append(executables, entry.Name())
		}
	}

	return executables, nil
}

func buildCompleter() Completer {
	completer := Completer{
		commands: []string{
			"echo",
			"exit",
		},
	}

	// executables
	path := os.Getenv("PATH")

	for dir := range strings.SplitSeq(path, string(os.PathListSeparator)) {
		executables, err := listExecutables(dir)
		if err != nil {
			continue
		}

		for _, executable := range executables {
			completer.commands = append(completer.commands, executable)
		}
	}
	
	return completer
}
