package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // next reading position in input
	ch           byte // current char

	// Both position and readPosition are to access characters in input by using
	// them as an index, ex.: l.input[l.readPosition]. These 2 “pointers” point into
	// our input string because we need to “peek” ahead to see what comes
	// next after the current character. readPosition always points to the “next”
	// input character. position points to input character corresponding to
	// the ch byte.
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Only for ASCII characters, not full UTF-8
func (l *Lexer) readChar() {

	// If hit end of input, set ch = 0 (NUL character).
	// So we've read nothing yet or it's EOF.
	if l.readPosition >= len(l.input) {
		l.ch = 0

	} else {
		l.ch = l.input[l.readPosition] // next character
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
