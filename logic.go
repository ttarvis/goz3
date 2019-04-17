package z3

import (
	"errors"
)

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

// Bool just wraps an AST node
// AST nodes are nodes in the SMT-Lib language
// We have to generate syntax of SMT in order to check it
// The logical methods below are not meant for non Bool sorts
type Bool struct {
	*AST
}

// NewBool returns a Bool const
// name can be either a 
func (ctx *Context) NewBool(name interface{}) (*Bool, error) {
	boolSort := ctx.MakeBool()
	var varSym *Symbol;
	switch v := name.(type) {
	case int:
		varSym = ctx.IntSymbol(v);	
	case string:
		varSym = ctx.StringSymbol(v);
	default:
		return nil, errors.New("Bool Const name type must be int or string");
	}

	boolConst := &Bool{
		ctx.Const(varSym, boolSort),
	}

	return boolConst, nil;
}

// Not wraps an AST in a Bool.
func (l *Bool) Not() *Bool {
	not := &Bool{
		&AST{
			Z3Context:	l.Z3Context,
			Z3AST:		C.Z3_mk_not(l.Z3Context, l.Z3AST),
		},
	}

	return not;
}

// Not creates an AST node representing not(ast)
// todo: bool
func (ast *AST) Not() *AST {
	not := &AST{
		Z3Context:	ast.Z3Context,
		Z3AST:		C.Z3_mk_not(ast.Z3Context, ast.Z3AST),
	}

	return not;
}

// And creates an AST node wrapped in a bool representing args[0] AND args... AND args[n-1]
// Be sure each AST added in this way is a part of the same context
func (l *Bool) And(r ...*Bool) *Bool {
	// make an array of C.Z3_ast of length of nodes + 1 to include the Bool used in method call
	args := make([]C.Z3_ast, len(r)+1);
	// add in the base Bool
	args[0] = l.Z3AST;
	for i, expr := range r {
		// todo: add check that context is the same in each stmt(node)
		// if the contexts are not the same, something bad will likely happen
		args[i+1] = expr.Z3AST;
	}

	val := &Bool{
		&AST{
			Z3Context:	l.Z3Context,
			Z3AST:		C.Z3_mk_and(l.Z3Context, C.uint(len(args)), &args[0]),
		},
	}

	return val;
}
 
// And creates an AST node representing args[0] AND args... AND args[n-1]
// Be sure each AST added in is a part of the same context
// ast and nodes should all be bools
// todo: bool
func (ast *AST) And(nodes...*AST) *AST {
	// make an array of Z3_ast of length of nodes + 1 to include *AST used in method call
	args := make([]C.Z3_ast, len(nodes)+1);
	// add in the base AST
	args[0] = ast.Z3AST;
	for i, node := range nodes {
		// todo: add check that context of each AST is the same
		// if the contexts are not the same, something bad will likely happen
		args[i+1] = node.Z3AST;
	}

	val := &AST{
		Z3Context:	ast.Z3Context,
		Z3AST:		C.Z3_mk_and(ast.Z3Context, C.uint(len(args)), &args[0]),
	}

	return val;
}

func (l *Bool) Or(r ...*Bool) *Bool {
	// make an array of C.Z3_ast of length of expr in r + 1 to include *Bool used in method cal
	args := make([]C.Z3_ast, len(r)+1);
	// add in the base Bool's AST
	args[0] = l.Z3AST;
	for i, expr := range r {
		// todo: add check that context of each AST is the same
		// if the contexts are not the same, something bad will likely happen
		args[i+1] = expr.Z3AST;
	}

	val := &Bool{
		&AST{
			Z3Context:	l.Z3Context,
			Z3AST:		C.Z3_mk_or(l.Z3Context, C.uint(len(args)), &args[0]),
		},
	}

	return val;
}

// Or creates an AST node representing args[0] OR ... OR args[n-1]
// todo: bool
func (ast *AST) Or(nodes...*AST) *AST {
	// make an array of Z3_ast of length of nodes + 1 to include *AST used in method call
	args := make([]C.Z3_ast, len(nodes)+1);
	// add in the base AST
	args[0] = ast.Z3AST;
	for i, node := range nodes {
		// todo: add check that context of each AST is the same
		// if the contexts are not the same, something bad will likely happen
		args[i+1] = node.Z3AST;
	}

	val := &AST{
		Z3Context:	ast.Z3Context,
		Z3AST:		C.Z3_mk_or(ast.Z3Context, C.uint(len(args)), &args[0]),
	}

	return val;
}

// Iff is for bools, same as IFF otherwise
func (t1 *Bool) Iff(t2 *Bool) *Bool {
	val := &Bool{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_iff(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return val;
}

// IFF creates an AST node representing t1 iff t2
// t1 and t2 must have the same context
// todo: bool
func (t1 *AST) IFF(t2 *AST) *AST {
	val := &AST{
		Z3Context:	t1.Z3Context,
		Z3AST:		C.Z3_mk_iff(t1.Z3Context, t1.Z3AST, t2.Z3AST),
	}

	return val;
}

// Eq returns a Bool that is true if l and r are equal
func (l *Bool) Eq(r *Bool) *Bool {
	val := &Bool{
		&AST{
			Z3Context:	l.Z3Context,
			Z3AST:		C.Z3_mk_eq(l.Z3Context, l.Z3AST, r.Z3AST),
		},
	}

	return val;
}

//func (l *Bool) Neq(r *Bool) *Bool

// Xor returns a Bool is true if l xor r
func (l *Bool) Xor(r *Bool) *Bool {
	val := &Bool{
		&AST{
			Z3Context:	l.Z3Context,
			Z3AST:		C.Z3_mk_xor(l.Z3Context, l.Z3AST, r.Z3AST),
		},
	}

	return val;
}

// Implies returns a Bool that is true if l implies r.
func (l *Bool) Implies(r *Bool) *Bool {
	val := &Bool{
		&AST{
			Z3Context:	l.Z3Context,
			Z3AST:		C.Z3_mk_implies(l.Z3Context, l.Z3AST, r.Z3AST),
		},
	}

	return val;
}
