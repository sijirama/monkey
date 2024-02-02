package lexer

import (
	"monko/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	test := []struct {
		expectedType    token.TokenType
		expextedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

    l := New(input)

    for i , tt := range test {
        tok := l.
    }

}
