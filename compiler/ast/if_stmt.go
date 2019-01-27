package ast

import "github.com/d5/tengo/compiler/source"

// IfStmt represents an if statement.
type IfStmt struct {
	IfPos source.Pos `json:"if_pos"`
	Init  Stmt       `json:"init"`
	Cond  Expr       `json:"cond"`
	Body  *BlockStmt `json:"body"`
	Else  Stmt       `json:"else"` // else branch; or nil
}

func (s *IfStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *IfStmt) Pos() source.Pos {
	return s.IfPos
}

// End returns the position of first character immediately after the node.
func (s *IfStmt) End() source.Pos {
	if s.Else != nil {
		return s.Else.End()
	}

	return s.Body.End()
}

func (s *IfStmt) String() string {
	var initStmt, elseStmt string
	if s.Init != nil {
		initStmt = s.Init.String() + "; "
	}
	if s.Else != nil {
		elseStmt = " else " + s.Else.String()
	}

	return "if " + initStmt + s.Cond.String() + " " + s.Body.String() + elseStmt
}
