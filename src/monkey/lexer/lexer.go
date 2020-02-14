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

	l.skipWhitespace()

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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			//The early exit here, our return tok statement, is necessary because when
			//calling readIdentifier(), we call readChar() repeatedly and advance our
			//readPosition and position fields past the last character of the current
			//identifier. So we don’t need the call to readChar() after the switch
			//statement again.
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

// If current character is a letter, read rest of the identifier/keyword
// till we hit a non-letter. Having read that identifier/keyword, find out if
// it's an identifier or keyword, to use correct token.TokenType.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// treat "_" as letter, so identifiers/keywords may have it, like "foo_bar"
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
