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
	l.skipWhiteSpace()

	var tok token.Token
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = token.NewToken(token.ASSIGN, '=')
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = token.NewToken(token.BANG, '!')
		}
	case ';':
		tok = token.NewToken(token.SEMICOLON, ';')
	case ',':
		tok = token.NewToken(token.COMMA, ',')
	case '(':
		tok = token.NewToken(token.LPAREN, '(')
	case ')':
		tok = token.NewToken(token.RPAREN, ')')
	case '{':
		tok = token.NewToken(token.LBRACE, '{')
	case '}':
		tok = token.NewToken(token.RBRACE, '}')
	case '+':
		tok = token.NewToken(token.PLUS, '+')
	case '-':
		tok = token.NewToken(token.MINUS, '-')

	case '*':
		tok = token.NewToken(token.ASTERISK, '*')
	case '/':
		tok = token.NewToken(token.SLASH, '/')
	case '<':
		tok = token.NewToken(token.LT, '<')
	case '>':
		tok = token.NewToken(token.GT, '>')

	case 0:
		tok = token.NewToken(token.EOF, 0)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return &tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return &tok
		} else {
			//パース出来なかった文字がくるところ
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return &tok
}

//改行コードや空白などでスキップするところで使うやーつ
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
		l.readChar()
	}
}
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 変数に使える文字を定義するところ
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
