package scanner

type TokenType int

//const (
//	LEFT_PAREN TokenType = iota
//	RIGHT_PAREN
//	LEFT_BRACE
//	RIGHT_BRACE
//	COMMA
//	DOT
//	MINUS
//	PLUS
//	SEMICOLON
//	SLASH
//	STAR
//	EQUAL
//	EQUAL_EQUAL
//	LESS
//	LESS_EQUAL
//	GREATER
//	GREATER_EQUAL
//	BANG
//	BANG_EQUAL
//	STRING
//	NUMBER
//	IDENTIFIER
//	EOF
//)
//
//func (t TokenType) String() string {
//	switch t {
//	case LEFT_PAREN:
//		return "LEFT_PAREN"
//	case RIGHT_PAREN:
//		return "RIGHT_PAREN"
//	case LEFT_BRACE:
//		return "LEFT_BRACE"
//	case RIGHT_BRACE:
//		return "RIGHT_BRACE"
//	case COMMA:
//		return "COMMA"
//	case PLUS:
//		return "PLUS"
//	case MINUS:
//		return "MINUS"
//	case SEMICOLON:
//		return "SEMICOLON"
//	case STAR:
//		return "STAR"
//	case DOT:
//		return "DOT"
//	case SLASH:
//		return "SLASH"
//	case EOF:
//		return "EOF"
//	case GREATER:
//		return "GREATER"
//	case GREATER_EQUAL:
//		return "GREATER_EQUAL"
//	case EQUAL:
//		return "EQUAL"
//	case EQUAL_EQUAL:
//		return "EQUAL_EQUAL"
//	case LESS:
//		return "LESS"
//	case LESS_EQUAL:
//		return "LESS_EQUAL"
//	case BANG_EQUAL:
//		return "BANG_EQUAL"
//	case BANG:
//		return "BANG"
//	case STRING:
//		return "STRING"
//	case NUMBER:
//		return "NUMBER"
//	case IDENTIFIER:
//		return "IDENTIFIER"
//	default:
//		return ""
//	}
//}
//

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	EQUAL
	EQUAL_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	BANG
	BANG_EQUAL
	STRING
	NUMBER
	IDENTIFIER
	AND
	CLASS
	ELSE
	FALSE
	FOR
	FUN
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

func (t TokenType) String() string {
	switch t {
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case COMMA:
		return "COMMA"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case SEMICOLON:
		return "SEMICOLON"
	case STAR:
		return "STAR"
	case DOT:
		return "DOT"
	case SLASH:
		return "SLASH"
	case EOF:
		return "EOF"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case EQUAL:
		return "EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case BANG:
		return "BANG"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
	case AND:
		return "AND"
	case CLASS:
		return "CLASS"
	case ELSE:
		return "ELSE"
	case FALSE:
		return "FALSE"
	case FOR:
		return "FOR"
	case FUN:
		return "FUN"
	case IF:
		return "IF"
	case NIL:
		return "NIL"
	case OR:
		return "OR"
	case PRINT:
		return "PRINT"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case TRUE:
		return "TRUE"
	case VAR:
		return "VAR"
	case WHILE:
		return "WHILE"

	default:
		return ""
	}
}

type Token struct {
	Type       TokenType
	Lexeme     string
	Literal    any
	LineNumber int
}

var Keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}
