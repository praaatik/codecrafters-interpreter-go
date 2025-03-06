package scanner

import (
	"fmt"
	"os"
	"unicode"
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
			s.line++ // increment the line after the comments are parsed?
		} else {
			s.addToken(Token{
				Type:       SLASH,
				Lexeme:     "/",
				Literal:    nil,
				LineNumber: s.line,
			})
		}
	case '"':
		s.scanString()
	case '\n':
		s.line++
		return
	case '\t':
		return
	case '\r':
		return
	case ' ':
		return

	// TODO: handle the default
	default:
		if unicode.IsNumber(rune(current)) {
			s.scanNumber()
		} else if s.isAlpha(rune(current)) {
			s.scanIdentifiers()
		} else {
			s.reportError(current)
		}
	}
}

func (s *Scanner) scanIdentifiers() {
	for s.isAlphanumeric(s.Peek()) {
		s.Advance()
	}

	var currentTokenType TokenType

	token, exists := Keywords[s.source[s.start:s.current]]
	if !exists {
		currentTokenType = IDENTIFIER
	} else {
		currentTokenType = token
	}

	s.addToken(Token{
		Type:       currentTokenType,
		Lexeme:     s.source[s.start:s.current],
		Literal:    nil,
		LineNumber: s.line,
	})
}

func (s *Scanner) isAlphanumeric(c byte) bool {
	return s.isAlpha(rune(c)) || unicode.IsNumber(rune(c))
}

func (s *Scanner) isAlpha(current rune) bool {
	return unicode.IsLetter(current) || current == '_'
}

func (s *Scanner) scanNumber() {
	integerPart := string(s.PeekPrev())
	decimalPart := ""

	for unicode.IsNumber(rune(s.Peek())) {
		integerPart = integerPart + string(s.Peek())
		s.Advance()
	}

	if string(s.Peek()) == "." && unicode.IsNumber(rune(s.PeekNext())) {
		s.Advance() // .

		for unicode.IsNumber(rune(s.Peek())) {
			decimalPart = decimalPart + string(s.Peek())
			s.Advance()
		}
	}

	literal := integerPart
	lexeme := integerPart
	if decimalPart != "" {
		lexeme += "."
		lexeme += decimalPart
	}

	isZero := true

	for _, value := range decimalPart {
		if value != '0' {
			isZero = false
		}
	}

	if isZero && decimalPart != "" {
		decimalPart = "0"
	}

	if decimalPart != "" {
		literal += "."
		literal += decimalPart
	} else {
		literal += ".0"
	}

	s.addToken(Token{
		Type:       NUMBER,
		Lexeme:     lexeme,
		Literal:    literal,
		LineNumber: s.line,
	})
}

func (s *Scanner) scanString() {
	startIndex := s.current

	//till not at the end AND not a double quote, continue
	for !s.isAtEnd() && s.source[s.current] != '"' {
		// continue even if on next line
		if s.source[s.current] == '\n' {
			s.line++
		}
		s.Advance()
	}

	// this is triggered only if the EOF is reached AND no double quotes were found
	// if double quotes were found, they'd be reached and handled in the previous for loop
	if s.isAtEnd() {
		_, _ = fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", s.line)
		s.hasError = true
		return
	}

	// consume the second double quote
	s.Advance()

	lexeme := s.source[s.start:s.current]         // include the double quotes
	literal := s.source[startIndex : s.current-1] // exclude the double quotes

	s.addToken(Token{
		Type:       STRING,
		Lexeme:     lexeme,
		Literal:    literal,
		LineNumber: s.line,
	})
}

func (s *Scanner) PeekPrev() byte {
	if s.current == 0 {
		return 0
	}
	return s.source[s.current-1]
}

func (s *Scanner) Peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) PeekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) reportError(c byte) {
	if c == 0 {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", s.line, c)
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

		switch token.Literal.(type) {
		case string:
			fmt.Println(token.Type.String(), lexeme, token.Literal)
		default:
			fmt.Println(token.Type.String() + " " + lexeme + " " + "null")
		}

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
