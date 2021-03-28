package lexer

import (
	"github.com/hideA88/monkey-lang/pkg/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(source string) *Lexer {
	l := &Lexer{
		input: source,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() *token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, '=')
	case ';':
		tok = token.NewToken(token.SEMICOLON, ';')
	case '(':
		tok = token.NewToken(token.LPAREN, '(')
	case ')':
		tok = token.NewToken(token.RPAREN, ')')
	case '{':
		tok = token.NewToken(token.LBRACE, '{')
	case '}':
		tok = token.NewToken(token.RBRACE, '}')
	case ',':
		tok = token.NewToken(token.COMMA, ',')
	case '+':
		tok = token.NewToken(token.PLUS, '+')
	case '-':
		tok = token.NewToken(token.MINUS, '-')
	case 0:
		tok = token.NewToken(token.EOF, 0)
	}

	l.readChar()
	return &tok
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
