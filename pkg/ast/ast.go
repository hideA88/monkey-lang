package ast

import "github.com/hideA88/monkey-lang/pkg/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token *token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {} //Expressionインターフェースとして扱えるようにするため
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct {
	Token *token.Token /// token.Let token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {} //Statementインターフェースとして扱えるようにするため
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
