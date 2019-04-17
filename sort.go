package z3


/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

// Sort wraps a Z3 sort type
type Sort struct {
	Ctx	C.Z3_context;
	Z3Sort 	C.Z3_sort;
}

// todo: maybe create a bool wrapper around sort so
// that bool is a type of sort
func (ctx *Context) MakeBool() *Sort {
	boolSort := &Sort {
		Ctx:	ctx.Z3Context,
		Z3Sort:	C.Z3_mk_bool_sort(ctx.Z3Context),
	}

	return boolSort;
}

func (ctx *Context) MakeInt() *Sort {
	intSort := &Sort {
		Ctx:	ctx.Z3Context,
		Z3Sort:	C.Z3_mk_int_sort(ctx.Z3Context),
	}

	return intSort;
}	

func (ctx *Context) MakeReal() *Sort {
	realSort := &Sort {
		Ctx:	ctx.Z3Context,
		Z3Sort:	C.Z3_mk_real_sort(ctx.Z3Context),
	}

	return realSort;
}

func (ctx *Context) MakeBV(sz uint) *Sort {
	bvSort := &Sort{
		Ctx:	ctx.Z3Context,
		Z3Sort:	C.Z3_mk_bv_sort(ctx.Z3Context, C.uint(sz)),
	}

	return bvSort;
}
