package main

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
	NUMBER
	IDENTIFIER
)

var reservedKeywords = []string{
	"and", "class", "else", "false", "true", "for", "fun",
	"if", "nil", "or", "print", "return", "super", "this", "var", "while",
}
