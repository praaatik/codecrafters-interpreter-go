package scanner

type TokenType int

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
	LESS_EQUAL
	GREATER_EQUAL
	BANG
	BANG_EQUAL
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
	case EQUAL:
		return "EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case BANG:
		return "BANG"

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
