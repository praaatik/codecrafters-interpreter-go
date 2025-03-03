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

// match checks if the current character matches the expected character.
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() || s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
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

	case '!':
		if s.match('=') {
			s.addToken(Token{
				Type:       BANG_EQUAL,
				Lexeme:     "!=",
				Literal:    nil,
				LineNumber: s.line,
			})
		} else {
			s.addToken(Token{
				Type:       BANG,
				Lexeme:     "!",
				Literal:    nil,
				LineNumber: s.line,
			})
		}

	case '=':
		if s.match('=') {
			s.addToken(Token{
				Type:       EQUAL_EQUAL,
				Lexeme:     "==",
				Literal:    nil,
				LineNumber: s.line,
			})
		} else {
			s.addToken(Token{
				Type:       EQUAL,
				Lexeme:     "=",
				Literal:    nil,
				LineNumber: s.line,
			})
		}

	case '<':
		if s.match('=') {
			s.addToken(Token{
				Type:       LESS_EQUAL,
				Lexeme:     "<=",
				Literal:    nil,
				LineNumber: s.line,
			})
		} else {
			s.addToken(Token{
				Type:       LESS,
				Lexeme:     "<",
				Literal:    nil,
				LineNumber: s.line,
			})
		}
	case '>':
		if s.match('=') {
			s.addToken(Token{
				Type:       GREATER_EQUAL,
				Lexeme:     ">=",
				Literal:    nil,
				LineNumber: s.line,
			})
		} else {
			s.addToken(Token{
				Type:       GREATER,
				Lexeme:     ">",
				Literal:    nil,
				LineNumber: s.line,
			})
		}
	case '/':
		if s.match('/') {
			for s.Advance() != '\n' && !s.isAtEnd() {
			}
		} else {
			s.addToken(Token{
				Type:       SLASH,
				Lexeme:     "/",
				Literal:    nil,
				LineNumber: s.line,
			})
		}

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

	_, _ = fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", c)
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
