package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // next reading position in input
	ch           byte // current char

	// Both position and readPosition access characters in input by using
	// them as an index. Ex: l.input[l.readPosition]. These 2 cursors point into
	// our input string because we must “peek” ahead to see what comes
	// next after the current character.
	// "readPosition" always points to the next input character.
	// "position" points to input character corresponding to the ch byte.
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Read 1 new character.
// "position": current index read. "readPosition": next index to be read.
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

// given input character, return its token
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
			// Calling readIdentifier() loops through readChar() repeatedly and advances our
			// readPosition and position fields past last character of current identifier.
			// Returns string literal that starts with l.ch.
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) // returns TokenType of literal
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

// If current character is a letter, loop through rest of the
// identifier/keyword till we hit a non-letter. Return substring.
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

//
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// like readIdentifier() but for digits
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
