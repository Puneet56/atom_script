package parser

import (
	"atom_script/ast"
	"atom_script/lexer"
	"testing"
)

func TestAtomStatement(t *testing.T) {
	input := `
	atom a = 1;
	atom b = 2;

	atom foobar = 123223;
	`

	// input := `
	// atom a  1;
	// atom  = 2;

	// atom 123223;
	// `

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

func TestProduceStatement(t *testing.T) {

	input := `
		produce 5;
		produce 10;
		produce 99922299;
	`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program statements does not contain 3 statement. got=%d\n",
			len(program.Statements),
		)
	}

	for _, stmt := range program.Statements {
		produceStmt, ok := stmt.(*ast.ProduceStatementStruct)

		if !ok {
			t.Errorf("statement not *ast.Statement. got=%T\n", stmt)
			continue
		}

		if produceStmt.TokenLiteral() != "produce" {
			t.Errorf("token literal of statement is not produce. got=%q", produceStmt.TokenLiteral())
		}

	}

}

func TestIdentiferExpression(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
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
