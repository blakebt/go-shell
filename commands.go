package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ShellState struct {
	pwd string
}

var state ShellState = ShellState{""}

func (s *ShellState) getState() string {
	return s.pwd
}

func (s *ShellState) setState(path string) {
	s.pwd = path
}

func executeCmd(cmd string, args []string) {
	switch cmd {
	case "echo":
		echo(args)
	case "pwd":
		printWorkingDir()
	case "cd":
		changeDir(args[0])
	case "date":
		getDate()
	case "quit":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println(cmd, "is not a valid command.")
		fmt.Println("Type 'help' for a list of available commands.")
	}
}

func changeDir(path string) {

	if newPath, err := filepath.Abs(path); err == nil {
		state.setState(newPath)
	}
}

func printHelp(args string) {
	commands := processFile("commands.txt")

	for _, c := range commands {
		fmt.Println(c)
	}
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func printWorkingDir() {
	fmt.Println(state.getState())
}

func getDate() {
	currTime := time.Now()

	fmt.Printf("%d-%v-%d %d:%d\n", currTime.Day(), currTime.Month(), currTime.Year(), currTime.Hour(), currTime.Minute())
}
