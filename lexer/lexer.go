package lexer

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
