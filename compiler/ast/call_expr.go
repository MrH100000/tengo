package ast

import (
	"strings"

	"github.com/d5/tengo/compiler/source"
)

// CallExpr represents a function call expression.
type CallExpr struct {
	Func   Expr       `json:"func"`
	LParen source.Pos `json:"lparen"`
	Args   []Expr     `json:"args"`
	RParen source.Pos `json:"rparen"`
}

func (e *CallExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CallExpr) Pos() source.Pos {
	return e.Func.Pos()
}

// End returns the position of first character immediately after the node.
func (e *CallExpr) End() source.Pos {
	return e.RParen + 1
}

func (e *CallExpr) String() string {
	var args []string
	for _, e := range e.Args {
		args = append(args, e.String())
	}

	return e.Func.String() + "(" + strings.Join(args, ", ") + ")"
}
