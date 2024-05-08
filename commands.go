package main

import (
	"bufio"
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
	case "rmdir":
		removeDir(args[0])
	case "touch":
		createFile(args[0])
	case "rm":
		deleteFile(args[0])
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
		log.Fatal(err)
	}

	return path
}

func changeDir(path string) {
	if newPath, err := filepath.Abs(path); err == nil {
		info, statErr := os.Stat(strings.TrimSpace(newPath))
		if statErr != nil {
			log.Fatal(statErr)
		}
		if info.IsDir() {
			newPath = strings.TrimSpace(newPath)
			err := os.Chdir(newPath)
			if err != nil {
				log.Fatal(err)
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

	fmt.Printf("%d-%v-%d %02d:%02d\n", currTime.Day(), currTime.Month(), currTime.Year(), currTime.Hour(), currTime.Minute())
}

func list() {
	pwd := getWorkingDir()

	entries, err := os.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	if len(entries) == 0 {
		return
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

func removeDir(path string) {
	path = strings.TrimSpace(path)

	entries, readErr := os.ReadDir(path)
	if readErr != nil {
		log.Fatal(readErr)
	}

	if len(entries) != 0 {
		fmt.Println("\n!!ALERT!!")
		fmt.Println("This directory is not empty. If you still wish to delete this directory and all its contents, please type 'y'. Otherwise type 'n'.")

		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'n' {
			return
		}
	}

	removeErr := os.RemoveAll(path)
	if removeErr != nil {
		log.Fatal(removeErr)
	}

	fmt.Printf("Directory successfully removed.\n")
}

func createFile(filename string) {
	filename = strings.TrimSpace(filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Println("File created successfully.")
}

func deleteFile(filename string) {
	filename = strings.TrimSpace(filename)

	err := os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %s removed successfully\n", filename)
}
