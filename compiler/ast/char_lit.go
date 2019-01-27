package ast

import "github.com/d5/tengo/compiler/source"

// CharLit represents a character literal.
type CharLit struct {
	Value    rune       `json:"value"`
	ValuePos source.Pos `json:"value_pos"`
	Literal  string     `json:"literal"`
}

func (e *CharLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CharLit) Pos() source.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *CharLit) End() source.Pos {
	return source.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *CharLit) String() string {
	return e.Literal
}
