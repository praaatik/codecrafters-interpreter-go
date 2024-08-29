package main

import (
	"fmt"
	"os"
)

// Lexeme -> Token

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	flag := false
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		s := NewScanner(fileContents)

		_, err = s.ScanTokens()
		if err != nil {
			flag = true
		}
	}

	if flag {
		fmt.Println("EOF  null")
		os.Exit(65)
	}

	fmt.Println("EOF  null")
}
