// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package syntax

import "fmt"

func walkStmts(sl StmtList, f func(Node) bool) {
	for _, s := range sl.Stmts {
		Walk(s, f)
	}
}

func walkWords(words []*Word, f func(Node) bool) {
	for _, w := range words {
		Walk(w, f)
	}
}

// Walk traverses an AST in depth-first order: It starts by calling
// f(node); node must not be nil. If f returns true, Walk invokes f
// recursively for each of the non-nil children of node, followed by
// f(nil).
func Walk(node Node, f func(Node) bool) {
	if !f(node) {
		return
	}

	switch x := node.(type) {
	case *File:
		walkStmts(x.StmtList, f)
	case *Stmt:
		if x.Cmd != nil {
			Walk(x.Cmd, f)
		}
		for _, r := range x.Redirs {
			Walk(r, f)
		}
	case *Assign:
		if x.Name != nil {
			Walk(x.Name, f)
		}
		if x.Value != nil {
			Walk(x.Value, f)
		}
		if x.Index != nil {
			Walk(x.Index, f)
		}
		if x.Array != nil {
			Walk(x.Array, f)
		}
	case *Redirect:
		if x.N != nil {
			Walk(x.N, f)
		}
		Walk(x.Word, f)
		if x.Hdoc != nil {
			Walk(x.Hdoc, f)
		}
	case *CallExpr:
		for _, a := range x.Assigns {
			Walk(a, f)
		}
		walkWords(x.Args, f)
	case *Subshell:
		walkStmts(x.StmtList, f)
	case *Block:
		walkStmts(x.StmtList, f)
	case *IfClause:
		walkStmts(x.Cond, f)
		walkStmts(x.Then, f)
		walkStmts(x.Else, f)
	case *WhileClause:
		walkStmts(x.Cond, f)
		walkStmts(x.Do, f)
	case *ForClause:
		Walk(x.Loop, f)
		walkStmts(x.Do, f)
	case *WordIter:
		Walk(x.Name, f)
		walkWords(x.Items, f)
	case *CStyleLoop:
		if x.Init != nil {
			Walk(x.Init, f)
		}
		if x.Cond != nil {
			Walk(x.Cond, f)
		}
		if x.Post != nil {
			Walk(x.Post, f)
		}
	case *BinaryCmd:
		Walk(x.X, f)
		Walk(x.Y, f)
	case *FuncDecl:
		Walk(x.Name, f)
		Walk(x.Body, f)
	case *Word:
		for _, wp := range x.Parts {
			Walk(wp, f)
		}
	case *Lit:
	case *SglQuoted:
	case *DblQuoted:
		for _, wp := range x.Parts {
			Walk(wp, f)
		}
	case *CmdSubst:
		walkStmts(x.StmtList, f)
	case *ParamExp:
		Walk(x.Param, f)
		if x.Index != nil {
			Walk(x.Index, f)
		}
		if x.Repl != nil {
			if x.Repl.Orig != nil {
				Walk(x.Repl.Orig, f)
			}
			if x.Repl.With != nil {
				Walk(x.Repl.With, f)
			}
		}
		if x.Exp != nil && x.Exp.Word != nil {
			Walk(x.Exp.Word, f)
		}
	case *ArithmExp:
		Walk(x.X, f)
	case *ArithmCmd:
		Walk(x.X, f)
	case *BinaryArithm:
		Walk(x.X, f)
		Walk(x.Y, f)
	case *BinaryTest:
		Walk(x.X, f)
		Walk(x.Y, f)
	case *UnaryArithm:
		Walk(x.X, f)
	case *UnaryTest:
		Walk(x.X, f)
	case *ParenArithm:
		Walk(x.X, f)
	case *ParenTest:
		Walk(x.X, f)
	case *CaseClause:
		Walk(x.Word, f)
		for _, ci := range x.Items {
			Walk(ci, f)
		}
	case *CaseItem:
		walkWords(x.Patterns, f)
		walkStmts(x.StmtList, f)
	case *TestClause:
		Walk(x.X, f)
	case *DeclClause:
		walkWords(x.Opts, f)
		for _, a := range x.Assigns {
			Walk(a, f)
		}
	case *ArrayExpr:
		for _, el := range x.Elems {
			Walk(el, f)
		}
	case *ArrayElem:
		if x.Index != nil {
			Walk(x.Index, f)
		}
		Walk(x.Value, f)
	case *ExtGlob:
		Walk(x.Pattern, f)
	case *ProcSubst:
		walkStmts(x.StmtList, f)
	case *TimeClause:
		if x.Stmt != nil {
			Walk(x.Stmt, f)
		}
	case *CoprocClause:
		if x.Name != nil {
			Walk(x.Name, f)
		}
		Walk(x.Stmt, f)
	case *LetClause:
		for _, expr := range x.Exprs {
			Walk(expr, f)
		}
	default:
		panic(fmt.Sprintf("syntax.Walk: unexpected node type %T", x))
	}

	f(nil)
}
