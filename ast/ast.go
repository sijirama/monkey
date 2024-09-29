package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//-----Ident------------------------------------------------------

type Identifier struct {
	Token token.Token //NOTE the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

//-----Let------------------------------------------------------

type LetStatement struct { // let x = 4;
	Token token.Token //the token.LET token
	Name  *Identifier //Name to hold the identifier of the binding
	Value Expression  //Value for the expression that produces the value
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

//-----------------------------------------------------------
