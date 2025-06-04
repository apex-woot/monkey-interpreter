package token

import "fmt"

type TokenType byte

func (t TokenType) String() string {
	if int(t) < len(tokenNames) {
		return tokenNames[t]
	}
	return fmt.Sprintf("UNKNOWN_TOKEN(%d)", t)
}

const (
	ILLEGAL TokenType = iota
	EOF
	// Identifiers
	IDENT
	INT

	// Delimiters
	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN

	// Operators
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	LT
	GT

	EQ
	NOT_EQ
)

// ORDER MATTERS AND SHOULD BE SAME AS IOTA CONSTANTS
var tokenNames = []string{
	"ILLEGAL",
	"EOF",
	// Identifiers
	"IDENT",
	"INT",

	// Delimiters
	"COMMA",
	"SEMICOLON",

	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",

	// Keywords
	"FUNCTION",
	"LET",
	"TRUE",
	"FALSE",
	"IF",
	"ELSE",
	"RETURN",

	// Operators
	"ASSIGN",
	"PLUS",
	"MINUS",
	"BANG",
	"ASTERISK",
	"SLASH",
	"LT",
	"GT",

	"EQ",
	"NOT_EQ",
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"eq":     EQ,
	"not_eq": NOT_EQ,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) String() string {
	return fmt.Sprintf("Token - Type=%q,Literal=%q", t.Type.String(), t.Literal)
}
