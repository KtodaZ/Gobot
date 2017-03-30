package main

import (
	"os/exec"
	"fmt"
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	isWindows bool   = true
	hardCodePath bool = false // Set to true to use the hardcoded basePath below
	basePath  string = "A:/GoPath/src/github.com/ktodaz/gobot" // Default path if not overriding
	cmd1, cmd2 *exec.Cmd
)

func main() {
	if !hardCodePath {
		basePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Path is " + basePath)
	}


	var firstFileName string
	fmt.Print("Enter first file's name: ")
	fmt.Scan(&firstFileName)

	var secondFileName string
	fmt.Print("Enter second file's name: ")
	fmt.Scan(&secondFileName)

	var  input string
	fmt.Print("Which file should go first (1 or 2): ")
	fmt.Scan(&input)
	shouldFirstFileGoFirst := input == "1"

	path1 := basePath + "/" + firstFileName
	path2 := basePath + "/" + secondFileName
	//fmt.Println("File 1 path is " + path1)
	//fmt.Println("File 2 path is " + path2)

	if isWindows {
		path1 += ".exe"
		path2 += ".exe"
	}

	if shouldFirstFileGoFirst {
		cmd1 = exec.Command(path1, "test", "true")
		cmd2 = exec.Command(path2, "test", "false")
	} else {
		cmd1 = exec.Command(path1, "test", "false")
		cmd2 = exec.Command(path2, "test", "true")
	}


	stdIn1, err := cmd1.StdinPipe()
	detectError(err, "error creating stdIn1")

	stdIn2, err := cmd2.StdinPipe()
	detectError(err, "error creating stdIn2")

	stdOut1, err := cmd1.StdoutPipe()
	detectError(err, "Error getting stdOut1 object")

	stdOut2, err := cmd2.StdoutPipe()
	detectError(err, "Error getting stdOut2 object")

	scanner1 := bufio.NewScanner(stdOut1)
	go func() {
		for scanner1.Scan() {
			fmt.Printf("file1: \t%s\n", scanner1.Text())
			if scanner1.Text() == "Won" {
				fmt.Println(firstFileName + " Won!")
				os.Exit(0)
			}
			if scanner1.Text() == "Lost" {
				fmt.Println(secondFileName + " Won!")
				os.Exit(0)
			}
			if scanner1.Text() != "Awaiting Input" && scanner1.Text() != "Input Received" {
				io.WriteString(stdIn2, scanner1.Text() + "\n")
			}
		}
	}()

	scanner2 := bufio.NewScanner(stdOut2)
	go func() {
		for scanner2.Scan() {
			fmt.Printf("file2: \t%s\n", scanner2.Text())
			if scanner2.Text() == "Won" {
				fmt.Println(secondFileName + " Won!")
				os.Exit(0)
			}
			if scanner2.Text() == "Lost" {
				fmt.Println(firstFileName + " Won!")
				os.Exit(0)
			}
			if scanner2.Text() != "Awaiting Input" && scanner2.Text() != "Input Received" {
				io.WriteString(stdIn1, scanner2.Text() + "\n")
			}
		}
	}()

	err = cmd1.Start()
	detectError(err, "Error starting cmd1")
	err = cmd2.Start()
	detectError(err, "Error starting cmd2")

	err = cmd1.Wait()
	detectError(err, "Error waiting for cmd1")
	err = cmd2.Wait()
	detectError(err, "Error waiting for cmd2")
}

func detectError(err error, msg string) {
	if err != nil {
		log.Panicf(msg, err)
	}
}
