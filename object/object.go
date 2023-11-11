package object

import (
	"atom_script/ast"
	"bytes"
	"fmt"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ       = "INTEGER"
	BOOLEAN_OBJ       = "BOOLEAN"
	NULL_OBJ          = "NULL"
	PRODUCE_VALUE_OBJ = "PRODUCE_VALUE_OBJ"
	ERROR_OBJ         = "ERROR"
	REACTION_OBJ      = "REACTION"
	STRING_OBJ        = "STRING"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type ProduceValue struct {
	Value Object
}

func (p *ProduceValue) Type() ObjectType { return PRODUCE_VALUE_OBJ }
func (p *ProduceValue) Inspect() string  { return p.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Reaction struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Reaction) Type() ObjectType { return REACTION_OBJ }

func (f *Reaction) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("reaction")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }
