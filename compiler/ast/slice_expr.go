package ast

import "github.com/d5/tengo/compiler/source"

// SliceExpr represents a slice expression.
type SliceExpr struct {
	Expr   Expr       `json:"expr"`
	LBrack source.Pos `json:"lbrack"`
	Low    Expr       `json:"low"`
	High   Expr       `json:"high"`
	RBrack source.Pos `json:"rbrack"`
}

func (e *SliceExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *SliceExpr) Pos() source.Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *SliceExpr) End() source.Pos {
	return e.RBrack + 1
}

func (e *SliceExpr) String() string {
	var low, high string
	if e.Low != nil {
		low = e.Low.String()
	}
	if e.High != nil {
		high = e.High.String()
	}

	return e.Expr.String() + "[" + low + ":" + high + "]"
}
