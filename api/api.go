package api

import (
	"atom_script/lexer"
	"atom_script/token"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/tokenize", handleTokenize)

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
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
