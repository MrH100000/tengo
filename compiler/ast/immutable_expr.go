package ast

import (
	"github.com/d5/tengo/compiler/source"
)

// ImmutableExpr represents an immutable expression
type ImmutableExpr struct {
	Expr         Expr       `json:"expr"`
	ImmutablePos source.Pos `json:"immutable_pos"`
	LParen       source.Pos `json:"lparen"`
	RParen       source.Pos `json:"rparen"`
}

func (e *ImmutableExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ImmutableExpr) Pos() source.Pos {
	return e.ImmutablePos
}

// End returns the position of first character immediately after the node.
func (e *ImmutableExpr) End() source.Pos {
	return e.RParen
}

func (e *ImmutableExpr) String() string {
	return "immutable(" + e.Expr.String() + ")"
}
