package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("\n----------------------------------------------------------------------")
	fmt.Println("\nWelcome to GoShell. Type 'help' for information on available commands.")
	fmt.Println("\n----------------------------------------------------------------------")

	for {

		if path, err := os.Getwd(); err == nil {
			state.setState(path)
		}

		fmt.Print(state.getState(), "$ ")
		in := bufio.NewReader(os.Stdin)

		line, err := in.ReadString('\n')
		// should never see this error because even spaces and newline characters
		// are valid
		if err != nil {
			panic(err)
		}

		line = strings.TrimRight(line, "\n")

		// split the user input on spaces to get the command and the arguments
		input := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' '
		})
		// the command is the first word entered and the args are everything
		// after the command
		cmd := strings.TrimSpace(input[0])
		args := input[1:]

		executeCmd(cmd, args)
	}

}

func processFile(fileName string) []string {
	// open the file
	readFile, err := os.Open(fileName)
	// check for an error when opening the file
	if err != nil {
		panic(err)
	}

	// read the file
	scanner := bufio.NewScanner(readFile)
	// split the file into lines
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	// scan each line into a slice
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	// no longer need the file, so close it
	readFile.Close()

	return fileLines
}
