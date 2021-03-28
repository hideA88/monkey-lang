package lexer

import "github.com/hideA88/monkey-lang/pkg/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(source string) *Lexer {
	return &Lexer{}
}

func (l *Lexer) NextToken() *token.Token {
	return &token.Token{
		Type:    token.EOF,
		Literal: "hogehoge",
	}
}
