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

// todo: can Int and Real be combined in to one type?
 
// Real wraps an AST node
type Real struct {
	*AST
}

// NewReal makes and returns an Real or an error
func (ctx *Context) NewReal(name interface{}) (*Real, error) {
	realSort := ctx.MakeReal();
	var varSym *Symbol;
	switch v := name.(type) {
	case int:
		varSym = ctx.IntSymbol(v);
	case string:
		varSym = ctx.StringSymbol(v);
	default:
		return nil, errors.New("Int Const name type must be int or string");
	}

	realConst := &Real{
		ctx.Const(varSym, realSort),
	}

	return realConst, nil;
}

// Add adds a number of Real types together
func (r *Real) Add(reals ...*Real) *Real {
	// make an array of C.Z3_ast of length of operands + 1 to include the Real used in method call
	args := make([]C.Z3_ast, len(reals)+1);
	// add in the first Real i
	args[0] = r.Z3AST;
	for n, operand := range reals {
		// todo: add check for same context
		args[n+1] = operand.Z3AST;
	}

	val := &Real{
		&AST{
			Z3Context:	r.Z3Context,
			Z3AST:		C.Z3_mk_add(r.Z3Context, C.uint(len(args)), &args[0]),
		},
	}

	return val;
}

// Mul returns the product of the args times i
// i * args[1] * ... * args[n-1]
// todo: ok, the naming of vars here is weird, reverting to l, r (left, right).
func (r *Real) Mul(reals ...*Real) *Real {
	args := make([]C.Z3_ast, len(reals)+1); // add one in for i
	args[0] = r.Z3AST;
	for i, n := range reals {
		args[i+1] = n.Z3AST;
	}

	val := &Real{
		&AST{
			Z3Context:	r.Z3Context,
			Z3AST:		C.Z3_mk_mul(r.Z3Context, C.uint(len(args)), &args[0]),
		},
	}

	return val;
}

// Sub returns l - arg[0] - ... - arg[n-1]
func (l *Real) Sub(rs ...*Real) *Real {
	args := make([]C.Z3_ast, len(rs)+1);
	args[0] = l.Z3AST;
	for i, r := range rs {
		args[i+1] = r.Z3AST;
	}

	val := &Real{
		&AST{
			Z3Context:	l.Z3Context,
			Z3AST:		C.Z3_mk_sub(l.Z3Context, C.uint(len(args)), &args[0]),
		},
	}

	return val;
}

// Div returns floor of n / d.
// If r == 0, result is ?
// todo: what is result
// this rounds toward -Inf whereas Go rounds toward 0.
func (n *Real) Div(d *Real) *Real {
	val := &Real{
		&AST{
			Z3Context:	n.Z3Context,
			Z3AST:		C.Z3_mk_div(n.Z3Context, n.Z3AST, d.Z3AST),
		},
	}

	return val;
}

// Minus returns -(n);
func (n *Real) Minus() *Real {
	val := &Real{
		&AST{
			Z3Context:	n.Z3Context,
			Z3AST:		C.Z3_mk_unary_minus(n.Z3Context, n.Z3AST),
		},
	}

	return val;
}

// Mod returns arg1 mod arg2
func (arg1 *Real) Mod(arg2 *Real) *Real {
	val := &Real{
		&AST{
			Z3Context:	arg1.Z3Context,
			Z3AST:		C.Z3_mk_mod(arg1.Z3Context, arg1.Z3AST, arg2.Z3AST),
		},
	}

	return val;
}

// Rem returns arg1 rem arg2
// todo: finish
func (arg1 *Real) Rem(arg2 *Real) *Real {
	return nil;
}

// Power returns arg1 ^ arg2
func (arg1 *Real) Power(arg2 *Real) *Real {
	return nil;
}

// LT returns node representing t1 < t2
func (t1 *Real) LT(t2 *Real) *Real {
	val := &Real{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_lt(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return val;
}

// LE returns node representing t1 <= t2
func (t1 *Real) LE(t2 *Real) *Real {
	val := &Real{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_le(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return val;
}

// GT returns node representing t1 > t2
func (t1 *Real) GT(t2 *Real) *Real {
	val := &Real{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_gt(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return val;
}

// GE returns node representing t1 >= t2
func (t1 *Real) GE(t2 *Real) *Real {
	val := &Real{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_ge(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return val;
}

// RealToInt returns Real value from Int
func (r *Real) RealToInt() *Real {
	val := &Real{
		&AST{
			Z3Context:	r.Z3Context,
			Z3AST:		C.Z3_mk_real2int(r.Z3Context, r.Z3AST),
		},
	}

	return val;
}
