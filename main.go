package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("\n----------------------------------------------------------------------")
	fmt.Println("\nWelcome to GoShell. Type 'help' for information on available commands.")
	fmt.Println("\n----------------------------------------------------------------------")

	for {
		path := getWorkingDir()

		fmt.Printf("%v$ ", path)

		line, err := in.ReadString('\n')
		// should never see this error because even spaces and newline characters
		// are valid
		if err != nil {
			log.Fatal(err)
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
