package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MetaCommandResult int
type PrepareResult int
type StatementType int
type Statement struct {
	sType StatementType
}

const (
	EXIT_SUCCESS                                        = 0
	EXIT_FAILURE                                        = 1
	META_COMMAND_SUCCESS              MetaCommandResult = 0
	META_COMMAND_UNRECOGNIZED_COMMAND MetaCommandResult = 1
	PREPARE_SUCCESS                   PrepareResult     = 0
	PREPARE_UNRECOGNIZED_STATEMENT    PrepareResult     = 1
	STATEMENT_INSERT                  StatementType     = 0
	STATEMENT_SELECT                  StatementType     = 1
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("db > ")
		command := readInput(reader)

		if strings.HasPrefix(command, ".") {
			switch doMetaCommand(command) {
			case META_COMMAND_SUCCESS:
				continue
			case META_COMMAND_UNRECOGNIZED_COMMAND:
				fmt.Printf("Unrecognized command %s.\n", command)
				continue
			}
		}

		stm := &Statement{}
		switch prepareStatement(command, stm) {
		case PREPARE_SUCCESS:
		case PREPARE_UNRECOGNIZED_STATEMENT:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", command)
			continue
		}

		executeStatement(stm)
		fmt.Println("Executed.")
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

func doMetaCommand(command string) MetaCommandResult {
	if command == ".exit" {
		os.Exit(EXIT_SUCCESS)
	}
	return META_COMMAND_UNRECOGNIZED_COMMAND

}

func prepareStatement(command string, statement *Statement) PrepareResult {
	if strings.HasPrefix(command, "select") {
		statement.sType = STATEMENT_SELECT
		return PREPARE_SUCCESS
	}
	if strings.HasPrefix(command, "insert") {
		statement.sType = STATEMENT_INSERT
		return PREPARE_SUCCESS
	}

	return PREPARE_UNRECOGNIZED_STATEMENT
}

func executeStatement(stm *Statement) {
	switch stm.sType {
	case STATEMENT_INSERT:
		fmt.Println("This is where we would do an insert")
	case STATEMENT_SELECT:
		fmt.Println("This is where we would do an select")
	}
}
