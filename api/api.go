package api

import (
	"atom_script/lexer"
	"atom_script/parser"
	"atom_script/token"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

	port := os.Getenv("PORT")

	if port == "" {
		port = "1323"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func handleTokenize(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	bodyBytes, err := io.ReadAll(body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	bodyString := string(bodyBytes)

	l := lexer.New(bodyString)

	tokens := make([]token.Token, 0)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tokens": tokens,
	})
}

func handleParsing(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	bodyBytes, err := io.ReadAll(body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	bodyString := string(bodyBytes)

	l := lexer.New(bodyString)

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
