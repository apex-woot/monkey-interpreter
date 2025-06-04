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
	// Identifiers and literals
	IDENT
	INT
	ASSIGN
	PLUS

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
)

var tokenNames = []string{
	"ILLEGAL",
	"EOF",
	// Identifiers and literals
	"IDENT",
	"INT",
	"ASSIGN",
	"PLUS",

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
}

type Token struct {
	Type    TokenType
	Literal string
}
