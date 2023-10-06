package lexer

import "atom_script/token"

// Lexer does the lexical analysis or tokenization of the input string.
type Lexer struct {
	input        string // input string to be tokenized
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination (ASCII only) complete list of ASCII codes: https://www.asciitable.com/
}

// New creates a new Lexer and initializes it with the input string.
// readChar() is called to initialize the Lexer's ch field.
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()
	return l
}

// readChar reads the next character and advances our position in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF (end of file) is reached, 0 is the ASCII code for the "NUL" character
	} else {
		l.ch = l.input[l.readPosition] // reading the next character
	}

	l.position = l.readPosition // position becomes the readPosition

	l.readPosition += 1 // readPosition advances by 1
}

// peekChar returns the next character without advancing our position in the input string.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition] // reading the next character
	}
}

// NextToken returns the next token and advances our position in the input string.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' { // if the next character is '='
			l.readChar() // read the next character
			tok = token.Token{
				Type:    token.EQ,
				Literal: "==",
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' { // if the next character is '='
			l.readChar() // read the next character
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: "!=",
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar() // read the next character
	return tok
}

// Skip the whitespace as we don't care about it.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar() // read the next character
	}
}

// just a helper function to create a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// read the identifier (atom, molecule, reaction, etc.)
func (l *Lexer) readIdentifier() string {
	position := l.position // position is the current position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position // position is the current position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
