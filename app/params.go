package main

import (
	"bufio"
	"github.com/google/shlex"
	"os"
)

func split(line string) []string {
	params, _ := shlex.Split(line)
	return params
}

func getParams() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		line := scanner.Text()
		return split(line), nil
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}
