package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

type AST struct {
	Z3Context	C.Z3_context;
	Z3AST		C.Z3_ast;		
}

// a value is basically anything that implements an AST
// I will just add a method to the AST itself to return the AST
// so that therefore any type I create counts as a value
// I can now add any type to certain solver methods like assert
// so that I don't need to create a separate assert for ASTs,
// Bools, etc.
// todo: I am not thrilled with this, it feels hackish.
// I guess it unifies all the types.
// an AST is a value, a Bool is a value, a {whatever} is a value
type value interface {
	GetAST()	C.Z3_ast;
}

// Const makes a "constant" from a symbol and it's sort(type).
// It returns an AST.  See details of ASTs and Consts in Z3 in
// comments above.
func (ctx *Context) Const(sym *Symbol, sort *Sort) *AST {
	ast := &AST{
		Z3Context:	ctx.Z3Context,
		Z3AST:		C.Z3_mk_const(ctx.Z3Context, sym.Z3Symbol, sort.Z3Sort),
	}

	return ast;
}

// implemented for the sake of the interface
func (ast *AST) GetAST() C.Z3_ast {
	return ast.Z3AST;
}

func (ast *AST) String() string {
	var s string;
	s = C.GoString(C.Z3_ast_to_string(ast.Z3Context, ast.Z3AST));
	return s
}
