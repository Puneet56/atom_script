package ast

import (
	"atom_script/token"
	"bytes"
)

// Node is the interface that all nodes in the AST implement.
// Node can be either a Statement or an Expression.
type Node interface {
	TokenLiteral() string
	String() string
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

type Identifier struct {
	token.Token // the token.IDENT token
	Value       string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
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

func (as *AtomStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.TokenLiteral() + " ")

	out.WriteString(as.Name.String() + " ")
	out.WriteString("=" + " ")

	if as.Value != nil {
		out.WriteString(as.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type ProduceStatementStruct struct {
	Token       token.Token
	ReturnValue Expression
}

func (ps *ProduceStatementStruct) statementNode() {}

func (ps *ProduceStatementStruct) TokenLiteral() string {
	return ps.Token.Literal
}

func (ps *ProduceStatementStruct) String() string {
	var out bytes.Buffer

	out.WriteString(ps.TokenLiteral() + " ")

	if ps.ReturnValue != nil {
		out.WriteString(ps.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
