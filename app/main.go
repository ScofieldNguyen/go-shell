package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Print("$ ")
	if scanner.Scan() {
		// Get command first
		command := scanner.Text()

		fmt.Printf("%s: command not found\n", command)
	}
}
