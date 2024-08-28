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
	EQUAL
	EQUAL_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	BANG
	BANG_EQUAL
	SLASH
	STRING
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

	//if s.source[s.current] == '\n' {
	//	s.line++
	//}

	return output
}

// ScanToken scans an individual lexeme and matches it the token types
func (s *Scanner) ScanToken() error {
	c := s.advance() // move forward

	isError := false
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
	case '!':
		if !s.match('=') {
			fmt.Println("BANG ! null")
			s.AddToken(BANG)
		} else {
			fmt.Println("BANG_EQUAL != null")
			s.AddToken(BANG_EQUAL)
		}
	case '=':
		if !s.match('=') {
			fmt.Println("EQUAL = null")
			s.AddToken(EQUAL)
		} else {
			fmt.Println("EQUAL_EQUAL == null")
			s.AddToken(EQUAL_EQUAL)
		}

	case '<':
		if !s.match('=') {
			fmt.Println("LESS < null")
			s.AddToken(LESS)
		} else {
			fmt.Println("LESS_EQUAL <= null")
			s.AddToken(LESS_EQUAL)
		}

	case '>':
		if !s.match('=') {
			fmt.Println("GREATER > null")
			s.AddToken(GREATER)
		} else {
			fmt.Println("GREATER_EQUAL >= null")
			s.AddToken(GREATER_EQUAL)
		}
	case '\n':
		s.line += 1
		//break
	case '\t':
	case ' ':
	case '/':
		if !s.match('/') {
			fmt.Println("SLASH / null")
			s.AddToken(SLASH)
		} else {
			for s.Peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		}

	case '"':
		err := s.string()
		if err != nil {
			return err
		}

	default:
		fmt.Fprintln(os.Stderr, fmt.Sprintf("[line %d] Error: Unexpected character: %c", s.line, c))
		isError = true
	}
	s.AddToken(EOF)
	if isError {
		return errors.New(fmt.Sprintf("[line 1] Error: Unexpected character: %c", c))
	}
	return nil
}

func (s *Scanner) Peek() byte {
	if !s.isAtEnd() {
		return s.source[s.current]
	}
	return 0
}

// ScanTokens scans the tokens one by one
func (s *Scanner) ScanTokens() ([]Token, error) {
	// check if we have reached the end of the current lexeme
	// if yes, set the start to the current token
	isError := false
	var err2 error
	for !s.isAtEnd() {
		s.start = s.current

		err := s.ScanToken()

		if err != nil {
			isError = true
			err2 = err
		}
	}

	if isError && err2 != nil {
		return s.tokens, errors.New(err2.Error())
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

// match matches the next character with the expected character and returns True if match, else false
func (s *Scanner) match(expectedCharacter byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expectedCharacter {
		return false
	}

	s.current += 1
	return true
}

func (s *Scanner) string() error {
	for s.Peek() != '"' && !s.isAtEnd() {
		if s.Peek() == '\n' {
			s.line += 1
		}

		s.advance()
	}

	if s.isAtEnd() {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("[line %d] Error: Unterminated string.", s.line))
		return errors.New("")
	}
	s.advance()

	stringValue := string(s.source[s.start+1 : s.current-1])
	s.AddLiteral(STRING, stringValue)

	fmt.Printf("STRING \"%s\" %s\n", stringValue, stringValue)

	return nil
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
