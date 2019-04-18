package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

// Model is a Z3 model wrapper
type Model struct {
	Z3Context	C.Z3_context;
	Z3Model		C.Z3_model;
}

// New Model create a fresh model
/* does not compile. Seems to be caused by Z3.
func (ctx *Context) NewModel() *Model {
	model := &Model{
		Z3Context:	ctx.Z3Context,
		Z3Model:	C.Z3_mk_model(ctx.Z3Context),
	}

	return model;
}
*/

// NumConstants returns a uint with the number of variables(constants) in the model
func (m *Model) NumConstants() uint {
	return uint(C.Z3_model_get_num_consts(m.Z3Context, m.Z3Model));
}

// GetConstInterp returns the interpretation (value) that Z3 assigned to each const (variable)
func (m *Model) GetConstInterp() map[string]*AST {
	consts := make(map[string]*AST);
	numConstants := m.NumConstants();

	for i := uint(0); i < numConstants; i++ {
		var cname	C.Z3_symbol;
		var cnst	C.Z3_func_decl;
		var a		C.Z3_ast;

		cnst	= C.Z3_model_get_const_decl(m.Z3Context, m.Z3Model, C.uint(i)); 
		cname	= C.Z3_get_decl_name(m.Z3Context, cnst);

		a = C.Z3_model_get_const_interp(m.Z3Context, m.Z3Model, cnst);

		name := C.GoString(C.Z3_get_symbol_string(m.Z3Context, cname));
		consts[name] = &AST{
					Z3Context:	m.Z3Context,
					Z3AST:		a,
				};
	}

	return consts;	
}
