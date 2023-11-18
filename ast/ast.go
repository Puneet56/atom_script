package ast

import (
	"atom_script/token"
	"bytes"
	"strings"
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

type MoleculeStatement struct {
	Token token.Token // the token.MOLECULE token
	Name  *Identifier
	Value Expression
}

func (ms *MoleculeStatement) statementNode() {}

func (ms *MoleculeStatement) TokenLiteral() string {
	return ms.Token.Literal
}

func (ms *MoleculeStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ms.TokenLiteral() + " ")

	out.WriteString(ms.Name.String() + " ")
	out.WriteString("=" + " ")

	if ms.Value != nil {
		out.WriteString(ms.Value.String())
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

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (inf *InfixExpression) expressionNode() {}

func (inf *InfixExpression) TokenLiteral() string {
	return inf.Token.Literal
}

func (inf *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(inf.Left.String())
	out.WriteString(" " + inf.Operator + " ")
	out.WriteString(inf.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type ReactionLiteral struct {
	Token      token.Token // The 'reaction' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *ReactionLiteral) expressionNode()      {}
func (fl *ReactionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *ReactionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or ReactionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type ArrayLiteral struct {
	Token    token.Token // The '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	Token token.Token // The '[' token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

type HashLiteral struct {
	Token token.Token // The '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }

func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}

	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
