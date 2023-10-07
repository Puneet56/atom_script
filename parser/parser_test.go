package parser

import (
	"atom_script/ast"
	"atom_script/lexer"
	"testing"
)

func TestAtomStatement(t *testing.T) {
	// input := `
	// atom a = 1;
	// atom b = 2;

	// atom foobar = 123223;
	// `

	input := `
	atom a  1;
	atom  = 2;

	atom 123223;
	`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
		{"b"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testAtomStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testAtomStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "atom" {
		t.Errorf("s.TokenLiteral not 'atom'. got=%q", s.TokenLiteral())
		return false
	}

	atomStmt, ok := s.(*ast.AtomStatement)

	if !ok {
		t.Errorf("s not *ast.AtomStatement. got=%T", s)
		return false
	}

	if atomStmt.Name.Value != name {
		t.Errorf("atomStmt.Name.Value not '%s'. got=%s", name, atomStmt.Name.Value)
		return false
	}

	if atomStmt.Name.TokenLiteral() != name {
		t.Errorf("atomStmt.Name.TokenLiteral() not '%s'. got=%s", name, atomStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
