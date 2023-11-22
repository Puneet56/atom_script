package api

import (
	"atom_script/evaluator"
	"atom_script/lexer"
	"atom_script/object"
	"atom_script/parser"
	"atom_script/token"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Code struct {
	Code string `json:"code"`
}

func logOutput() io.Writer {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return os.Stdout
	}

	return file
}

func Init() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
			Output: logOutput(),
		},
	))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	e.POST("/api/tokenize", handleTokenize)
	e.POST("/api/ast", handleGenerateAst)
	e.POST("/api/parse", handleParsing)
	e.POST("/api/eval", handleEval)
	e.POST("/api/repl", handleRepl)

	port := os.Getenv("PORT")

	if port == "" {
		port = "1323"
	}

	e.Logger.Printf("Starting server started on port %s", port)

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))
}

func handleTokenize(c echo.Context) error {
	var body Code

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	codeString := body.Code

	l := lexer.New(codeString)

	tokens := make([]token.Token, 0)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tokens": tokens,
	})
}

func handleGenerateAst(c echo.Context) error {
	var body Code

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	codeString := body.Code

	l := lexer.New(codeString)

	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": p.Errors(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"ast": program,
	})
}

func handleParsing(c echo.Context) error {
	var body Code

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	codeString := body.Code

	l := lexer.New(codeString)

	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": p.Errors(),
		})
	}

	resp := make([]string, 0)

	for _, stmt := range program.Statements {
		if stmt.String() == "" || stmt.String() == " " {
			continue
		}

		resp = append(resp, stmt.String())
	}

	return c.JSON(http.StatusOK, resp)
}

var env = object.NewEnvironment()

func handleEval(c echo.Context) error {
	var body Code

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	codeString := body.Code

	l := lexer.New(codeString)

	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": p.Errors(),
		})
	}

	response := make([]string, 0)

	for _, stmt := range program.Statements {
		evaluated := evaluator.Eval(stmt, env)

		if evaluated == nil {
			response = append(response, "null")
			continue
		}

		response = append(response, evaluated.Inspect())
	}

	return c.JSON(http.StatusOK, response)
}

type CodeBlock struct {
	Code       string `json:"code"`
	IsExecuted bool   `json:"isExecuted"`
}

type ReplCode struct {
	Code []CodeBlock `json:"code"`
}

func handleRepl(c echo.Context) error {
	var body ReplCode

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	response := make([]string, 0)

	env := object.NewEnvironment()
	for _, codeBlock := range body.Code {

		codeString := codeBlock.Code

		l := lexer.New(codeString)

		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"errors": p.Errors(),
			})
		}

		for _, stmt := range program.Statements {
			evaluated := evaluator.Eval(stmt, env)

			if evaluated == nil {
				continue
			}

			if !codeBlock.IsExecuted {
				response = append(response, evaluated.Inspect())
			}
		}
	}

	return c.JSON(http.StatusOK, response)
}
