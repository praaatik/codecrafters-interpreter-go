package main

import (
	"errors"
	"fmt"
	"os"
)

// Lexeme -> Token

type TokenType int

const (
	EOF TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	STAR
	DOT
	COMMA
	PLUS
	MINUS
	SEMICOLON
)

type Token struct {
	Type    TokenType   // Type would classify each lexeme
	Lexeme  string      // each word is a lexeme in the code
	Literal interface{} // literal values for strings and numbers
	Line    int         // store the line number to get the location information
}

type Scanner struct {
	source  []byte // raw code which is being read
	tokens  []Token
	start   int // index the start of the string
	current int // index the current character under examination
	line    int // track the current line for error handling
}

func NewScanner(source []byte) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) AddLiteral(tokenType TokenType, literal interface{}) {
	tt := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, Token{tokenType, tt, literal, s.line})
}

// AddToken calls the AddLiteral to add the literal type
func (s *Scanner) AddToken(tokenType TokenType) {
	s.AddLiteral(tokenType, nil)
}

// isAtEnd returns if the scanner has reached the end of the current lexeme
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// advance will move the current pointer to the next pointer in the source
func (s *Scanner) advance() byte {
	output := s.source[s.current]
	s.current++

	return output
}

// ScanToken scans an individual lexeme and matches it the token types
func (s *Scanner) ScanToken() error {
	c := s.advance() // move forward

	switch c {
	case '(':
		fmt.Println("LEFT_PAREN ( null")
		s.AddToken(LEFT_PAREN)
	case ')':
		fmt.Println("RIGHT_PAREN ) null")
		s.AddToken(RIGHT_PAREN)
	case '{':
		fmt.Println("LEFT_BRACE { null")
		s.AddToken(LEFT_BRACE)
	case '}':
		fmt.Println("RIGHT_BRACE } null")
		s.AddToken(RIGHT_BRACE)
	case '*':
		fmt.Println("STAR * null")
		s.AddToken(STAR)
	case '+':
		fmt.Println("PLUS + null")
		s.AddToken(PLUS)
	case '-':
		fmt.Println("MINUS - null")
		s.AddToken(MINUS)
	case ',':
		fmt.Println("COMMA , null")
		s.AddToken(COMMA)
	case '.':
		fmt.Println("DOT . null")
		s.AddToken(DOT)
	case ';':
		fmt.Println("SEMICOLON ; null")
		s.AddToken(SEMICOLON)
	default:
		fmt.Fprintln(os.Stderr, fmt.Sprintf("[line 1] Error: Unexpected character: %c", c))
	}
	s.AddToken(EOF)
	return nil
}

// ScanTokens scans the tokens one by one
func (s *Scanner) ScanTokens() ([]Token, error) {
	// check if we have reached the end of the current lexeme
	// if yes, set the start to the current token
	isError := false
	var err error
	for !s.isAtEnd() {
		s.start = s.current

		err = s.ScanToken()
		if err != nil {
			isError = true
		}
	}

	if isError && err != nil {
		return s.tokens, errors.New(err.Error())
	}

	// append the EOF token for the end of the file
	s.tokens = append(s.tokens, Token{
		Type:    EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    s.line,
	})

	return s.tokens, nil
}

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

	// Uncomment this block to pass the first stage
	//
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		s := NewScanner(fileContents)
		_, _ = s.ScanTokens()
	}
	fmt.Println("EOF  null")
}
