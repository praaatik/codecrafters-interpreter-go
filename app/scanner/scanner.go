package scanner

import (
	"fmt"
	"os"
)

type Scanner struct {
	source   string  // source code
	tokens   []Token // list of Tokens which were found
	start    int     // start of the string
	current  int     // current character under examination
	line     int     // current line under examincation
	hasError bool    // hasError flag is set if there are any errors during scanning
}

func (s *Scanner) Advance() byte {
	if s.isAtEnd() {
		return 0
	}

	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) scanToken() {
	current := s.Advance()

	switch current {
	case '(':
		s.addToken(Token{
			Type:       LEFT_PAREN,
			Lexeme:     "(",
			Literal:    nil,
			LineNumber: s.line,
		})
	case ')':
		s.addToken(Token{
			Type:       RIGHT_PAREN,
			Lexeme:     ")",
			Literal:    nil,
			LineNumber: 0,
		})
	case '{':
		s.addToken(Token{
			Type:       LEFT_BRACE,
			Lexeme:     "{",
			Literal:    nil,
			LineNumber: 0,
		})
	case '}':
		s.addToken(Token{
			Type:       RIGHT_BRACE,
			Lexeme:     "}",
			Literal:    nil,
			LineNumber: 0,
		})
	case '*':
		s.addToken(Token{
			Type:       STAR,
			Lexeme:     "*",
			Literal:    nil,
			LineNumber: 0,
		})
	case '.':
		s.addToken(Token{
			Type:       DOT,
			Lexeme:     ".",
			Literal:    nil,
			LineNumber: 0,
		})
	case '+':
		s.addToken(Token{
			Type:       PLUS,
			Lexeme:     "+",
			Literal:    nil,
			LineNumber: 0,
		})
	case '-':
		s.addToken(Token{
			Type:       MINUS,
			Lexeme:     "-",
			Literal:    nil,
			LineNumber: 0,
		})
	case ',':
		s.addToken(Token{
			Type:       COMMA,
			Lexeme:     ",",
			Literal:    nil,
			LineNumber: 0,
		})
	case ';':
		s.addToken(Token{
			Type:       SEMICOLON,
			Lexeme:     ";",
			Literal:    nil,
			LineNumber: 0,
		})
	case '\n':
		return

	// TODO: handle the default
	default:
		s.reportError(current)
	}
}

func (s *Scanner) reportError(c byte) {
	if c == 0 {
		return
	}

	fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", c)
	s.hasError = true
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.addToken(Token{
		Type:       EOF,
		Lexeme:     "",
		Literal:    nil,
		LineNumber: s.line,
	})
	return s.tokens
}

func (s *Scanner) addToken(token Token) {
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) PrintOutput() {
	for _, token := range s.tokens {
		var lexeme string
		if token.Type == EOF {
			lexeme = ""
		} else {
			lexeme = token.Lexeme
		}

		fmt.Println(token.Type.String() + " " + lexeme + " " + "null")
	}

	if s.hasError {
		os.Exit(65)
	}
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  nil,
		start:   0,
		current: 0,
		line:    1,
	}
}
