package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EXIT_SUCCESS = iota
	EXIT_FAILURE
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("db > ")
		command := readInput(reader)

		if command == ".exit" {
			os.Exit(EXIT_SUCCESS)
		} else {
			fmt.Printf("Unrecognized command %s.\n", command)
		}
	}
}

func readInput(reader *bufio.Reader) string {
	command, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	command = strings.TrimRight(command, "\n")
	if len(command) == 0 {
		fmt.Println("Error reading input")
		os.Exit(EXIT_FAILURE)
	}

	return command
}
