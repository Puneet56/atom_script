package ast

import "atom_script/token"

// Node is the interface that all nodes in the AST implement.
// Node can be either a Statement or an Expression.
type Node interface {
	TokenLiteral() string
}

// Statement is the interface that all statement nodes in the AST implement.
// Statement nodes do not produce a value.
type Statement interface {
	Node
	statementNode()
}

// Expression is the interface that all expression nodes in the AST implement.
// Expression nodes produce a value.
type Expression interface {
	Node
	expressionNode()
}

// Statement nodes in the AST
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

type AtomStatement struct {
	Token token.Token // the token.ATOM token
	Name  *Identifier
	Value Expression
}

func (as *AtomStatement) statementNode() {}

func (as *AtomStatement) TokenLiteral() string {
	return as.Token.Literal
}

type Identifier struct {
	token.Token // the token.IDENT token
	Value       string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
