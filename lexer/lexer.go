package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current reading position in input (after current char)
	readPosition int  /* current position in input (points to current char) readPosition always points to the “next” character in the input */
	char         byte // current char under examination, position points to the character in the input that corresponds to the ch byte.
}

// INFO: return a new lexer type with string input placed
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	//NOTE: i think about this but won't a map be faster for this, instead of a switch statement, at least for the operators
	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

// INFO: return a token type
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// INFO: The purpose of readChar is to give us the next character and advance our position in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // have we reached the end of the input ?
		l.char = 0 // if we have set l.char to 0, which is the ASCII code for the "NUL"
	} else {
		l.char = l.input[l.readPosition] //set l.char to the next character by accessing l.input[l.readPosition]
	}
	l.position = l.readPosition // update the l.position
	l.readPosition += 1         //increment readPosition to the next character

	//That way, l.readPosition always points to the next position where we’re going to read from next and
	//l.position always points to the position where we last read
}

// NOTE: read identifier
func (l *Lexer) readIdentifier() string {
	position := l.position // get the current position
	for isLetter(l.char) { // loop aslong as we encounter characters
		l.readChar() // keep getting the next char
	}
	return l.input[position:l.position] // creturn slice of last char stop to next space
}

// NOTE: self-explanatory
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}
