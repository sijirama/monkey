package lexer

import (
	"monko/token"
)

type Lexer struct {
	input        string
	position     int  //INFO: current position in input (points to current char)
	readPosition int  //INFO: current reading position in input (after current char)
	ch           byte //INFO: current char under examination
}

/*
INFO:
Most of the fields in Lexer are pretty self-explanatory. The ones that might cause some confusion
right now are position and readPosition. Both will be used to access characters in input by
using them as an index, e.g.: l.input[l.readPosition]. The reason for these two “pointers”
pointing into our input string is the fact that we will need to be able to “peek” further into
the input and look after the current character to see what comes up next. readPosition always
points to the “next” character in the input. position points to the character in the input that
corresponds to the ch byte.
*/

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//INFO: The purpose of readChar is to give us the next character and advance our position in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
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
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			return newToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok

}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

//INFO: this is letter function is responsible for allowing the characters we include in our identifiers e.g we are cheking for '_', so identifiers such as foo_bar will be allowed
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

//NOTE: to skip whitespace
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
