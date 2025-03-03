package main

import (
	"fmt"
	"os"

	scanner2 "github.com/codecrafters-io/interpreter-starter-go/app/scanner"
)

func main() {
	if len(os.Args) < 3 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		_, _ = fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	source := string(fileContents)

	scanner := scanner2.NewScanner(source)
	_ = scanner.ScanTokens()
	scanner.PrintOutput()
}
