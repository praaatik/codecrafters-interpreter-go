package main

type TokenType int

type Token struct {
	Type    TokenType   // Type would classify each lexeme
	Lexeme  string      // each word is a lexeme in the code
	Literal interface{} // literal values for strings and numbers
	Line    int         // store the line number to get the location information
}
