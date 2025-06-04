package lexer

import (
	"fmt"

	"github.com/apex-woot/monkey-interpreter/dlog"
	"github.com/apex-woot/monkey-interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) String() string {
	return fmt.Sprintf("Lexer[position=%d,readPosition=%d,ch=%q\n", l.position, l.readPosition, string(l.ch))
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	dlog.Debug.Printf("[lexer] created new lexer using \"\n%s\n\"\n", l.input)
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.eatWhitespace()
	dlog.Debug.Printf("[lexer] eat whitespace -  %s\n", &tok)
	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case 0:
		tok = newToken(token.EOF, l.ch)
	default:
		if isLetter(l.ch) {
			dlog.Debug.Printf("[lexer] %q is a letter\n", l.ch)
			tok.Literal = l.readLiteral(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			dlog.Debug.Printf("[lexer] %q is a digit\n", l.ch)
			tok.Literal = l.readLiteral(isDigit)
			tok.Type = token.INT
			return tok
		} else {
			dlog.Debug.Printf("[lexer] %q is unknown\n", l.ch)
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	dlog.Debug.Printf("l.ch - %q\n", l.ch)
	dlog.Debug.Printf("[lexer] returning %s\n", &tok)
	l.readChar()
	return tok
}

func (l *Lexer) readLiteral(condition func(byte) bool) string {
	position := l.position
	for condition(l.ch) {
		l.readChar()
	}
	literal := l.input[position:l.position]
	dlog.Debug.Printf("[lexer] read literal %s\n", literal)
	return literal
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) eatWhitespace() {
	dlog.Debug.Printf("[lexer] eating whitespace %s\n", l)
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tok token.TokenType, ch byte) token.Token {
	literal := string(ch)
	if ch == 0 {
		literal = ""
	}

	t := token.Token{Type: tok, Literal: literal}
	return t
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
