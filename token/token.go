package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	ATOM     = "ATOM"
	MOLECULE = "MOLECULE"
	REACTION = "REACTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	PROODUCE = "PRODUCE"
)

// Instead of let, const and fn we are using ATOM, MOLECULE and REACTION. We are also using PRODUCE instead of return.

var keywords = map[string]TokenType{
	"atom":     ATOM,
	"molecule": MOLECULE,
	"reaction": REACTION,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"produce":  PROODUCE,
}

// LookupIdent checks the keywords table to see whether the given identifier is in fact a keyword.
// If it is, it returns the keyword's TokenType constant.
// If it isn't, it assumes we are dealing with a user-defined identifier and returns the token.IDENT constant instead.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
