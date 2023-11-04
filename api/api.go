package api

import (
	"atom_script/evaluator"
	"atom_script/lexer"
	"atom_script/parser"
	"atom_script/token"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Code struct {
	Code string `json:"code"`
}

func Init() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		},
	))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/api/tokenize", handleTokenize)
	e.POST("/api/parse", handleParsing)
	e.POST("/api/eval", handleEval)
	e.POST("/api/repl", handleRepl)

	port := os.Getenv("PORT")

	if port == "" {
		port = "1323"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
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
		evaluated := evaluator.Eval(stmt)

		if evaluated == nil {
			response = append(response, "null")
		}

		response = append(response, evaluated.Inspect())
	}

	return c.JSON(http.StatusOK, response)
}

func handleRepl(c echo.Context) error {
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

	evaluated := evaluator.Eval(program)

	if evaluated == nil {
		return c.JSON(http.StatusOK, "null")
	}

	return c.JSON(http.StatusOK, evaluated.Inspect())
}
