package z3

import (
	"unsafe"
)

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

type Symbol struct {
	Z3Context	C.Z3_context;
	Z3Symbol	C.Z3_symbol;
}

func (ctx *Context) IntSymbol(i int) *Symbol {
	cInt := C.int(i);
	defer C.free(unsafe.Pointer(&cInt));

	return &Symbol{
		Z3Context:	ctx.Z3Context,
		Z3Symbol:	C.Z3_mk_int_symbol(ctx.Z3Context, cInt),
	}
}

func (ctx *Context) StringSymbol(s string) *Symbol {
	cname := C.CString(s);
	defer C.free(unsafe.Pointer(cname));
	sym := &Symbol{
		Z3Context:	ctx.Z3Context,
		Z3Symbol:	C.Z3_mk_string_symbol(ctx.Z3Context, cname),
	}

	return sym;
}
