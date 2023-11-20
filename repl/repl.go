package repl

import (
	"atom_script/evaluator"
	"atom_script/lexer"
	"atom_script/object"
	"atom_script/parser"
	"bufio"
	"fmt"
	"os"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Atom Script! Feel free to type in commands")
	fmt.Print(">> ")

	env := object.NewEnvironment()

	for scanner.Scan() {
		l := lexer.New(scanner.Text())

		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				fmt.Println(err)
			}
		}

		for _, stmt := range program.Statements {
			evaluated := evaluator.Eval(stmt, env)

			if evaluated != nil {
				fmt.Println(evaluated.Inspect())
			}

			fmt.Print(">> ")
		}
	}
}
