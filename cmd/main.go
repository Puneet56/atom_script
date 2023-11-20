package main

import (
	"atom_script/api"
	"atom_script/evaluator"
	"atom_script/lexer"
	"atom_script/object"
	"atom_script/parser"
	"atom_script/repl"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		repl.Start()
		return
	}

	switch args[0] {
	case "--api":
		fmt.Println("Starting API server...")
		api.Init()
	case "--file":
		if len(args) < 2 {
			fmt.Println("Please provide a file to run")
			return
		}

		fmt.Println("Running file...")
		evalFile(args[1])

	default:
		repl.Start()
	}
}

func evalFile(file string) {
	bytes, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("Error reading file, please check if the file exists")
		return
	}

	l := lexer.New(string(bytes))

	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			fmt.Println(err)
		}
	}

	env := object.NewEnvironment()

	for _, stmt := range program.Statements {
		evaluated := evaluator.Eval(stmt, env)

		if evaluated != nil {
			fmt.Println(evaluated.Inspect())
		}
	}
}
