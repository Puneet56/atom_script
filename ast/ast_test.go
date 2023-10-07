package ast

import (
	"atom_script/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&AtomStatement{
				Token: token.Token{
					Type:    token.ATOM,
					Literal: "atom",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "atom myVar = anotherVar;" {
		t.Errorf("program.String() is wrong. got=%q", program.String())
	}
}
