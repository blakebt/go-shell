package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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
	case "help":
		printHelp()
	case "mkdir":
		mkDir(args[0])
	case "ls":
		list()
	case "quit":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println(cmd, "is not a valid command.")
		fmt.Println("Type 'help' for a list of available commands.")
	}
}

func getWorkingDir() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path
}

func changeDir(path string) {
	if newPath, err := filepath.Abs(path); err == nil {
		info, statErr := os.Stat(strings.TrimSpace(newPath))
		if statErr != nil {
			panic(statErr)
		}
		if info.IsDir() {
			newPath = strings.TrimSpace(newPath)
			err := os.Chdir(newPath)
			if err != nil {
				panic(err)
			}
		}
	}
}

func mkDir(path string) {
	path = strings.TrimSpace(path)

	err := os.MkdirAll(path, 0750)

	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func printHelp() {
	commands := processFile("commands.txt")

	for _, c := range commands {
		fmt.Println(c)
	}
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func printWorkingDir() {
	fmt.Println(getWorkingDir())
}

func getDate() {
	currTime := time.Now()

	fmt.Printf("%d-%v-%d %d:%d\n", currTime.Day(), currTime.Month(), currTime.Year(), currTime.Hour(), currTime.Minute())
}

func list() {
	pwd := getWorkingDir()

	entries, err := os.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%-15v%-15v%v\n", "Mode", "Length", "Name")
	fmt.Printf("%-15v%-15v%v\n", "----", "------", "----")
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			log.Fatal(err)
		}

		fileType := "----"
		if info.Mode().IsDir() {
			fileType = "dir"
		}
		fmt.Printf("%-15v%-15v%v \n", fileType, info.Size(), e.Name())
	}
	fmt.Println()
}
