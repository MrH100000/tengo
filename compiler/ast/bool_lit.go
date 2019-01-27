package ast

import "github.com/d5/tengo/compiler/source"

// BoolLit represents a boolean literal.
type BoolLit struct {
	Value    bool       `json:"value"`
	ValuePos source.Pos `json:"value_pos"`
	Literal  string     `json:"literal"`
}

func (e *BoolLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BoolLit) Pos() source.Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *BoolLit) End() source.Pos {
	return source.Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *BoolLit) String() string {
	return e.Literal
}
