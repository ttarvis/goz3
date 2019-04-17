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

// Int wraps an AST node
type Int struct {
	*AST
}

// NewInt makes and returns an Int or an error
func (ctx *Context) NewInt(name interface{}) (*Int, error) {
	intSort := ctx.MakeInt();
	var varSym *Symbol;
	switch v := name.(type) {
	case int:
		varSym = ctx.IntSymbol(v);
	case string:
		varSym = ctx.StringSymbol(v);
	default:
		return nil, errors.New("Int Const name type must be int or string");
	}

	intConst := &Int{
		ctx.Const(varSym, intSort),
	}

	return intConst, nil;
}

// Add adds a number of Int types together
func (i *Int) Add(ints ...*Int) *Int {
	// make an array of C.Z3_ast of length of operands + 1 to include the Int used in method call
	args := make([]C.Z3_ast, len(ints)+1);
	// add in the first Int i
	args[0] = i.Z3AST;
	for n, operand := range ints {
		// todo: add check for same context
		args[n+1] = operand.Z3AST;
	}

	val := &Int{
		&AST{
			Z3Context:	i.Z3Context,
			Z3AST:		C.Z3_mk_add(i.Z3Context, C.uint(len(args)), &args[0]),
		},
	}

	return val;
}

// Mul returns the product of the args times i
// i * args[1] * ... * args[n-1]
// todo: ok, the naming of vars here is weird, reverting to l, r (left, right).
func (I *Int) Mul(ints ...*Int) *Int {
	args := make([]C.Z3_ast, len(ints)+1); // add one in for i
	args[0] = I.Z3AST;
	for i, n := range ints {
		args[i+1] = n.Z3AST;
	}

	val := &Int{
		&AST{
			Z3Context:	I.Z3Context,
			Z3AST:		C.Z3_mk_mul(I.Z3Context, C.uint(len(args)), &args[0]),
		},
	}

	return val;
}

// Sub returns l - arg[0] - ... - arg[n-1]
func (l *Int) Sub(rs ...*Int) *Int {
	args := make([]C.Z3_ast, len(rs)+1);
	args[0] = l.Z3AST;
	for i, r := range rs {
		args[i+1] = r.Z3AST;
	}

	val := &Int{
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
func (n *Int) Div(d *Int) *Int {
	val := &Int{
		&AST{
			Z3Context:	n.Z3Context,
			Z3AST:		C.Z3_mk_div(n.Z3Context, n.Z3AST, d.Z3AST),
		},
	}

	return val;
}

// Minus returns -(n);
func (n *Int) Minus() *Int {
	val := &Int{
		&AST{
			Z3Context:	n.Z3Context,
			Z3AST:		C.Z3_mk_unary_minus(n.Z3Context, n.Z3AST),
		},
	}

	return val;
}

// Mod returns arg1 mod arg2
func (arg1 *Int) Mod(arg2 *Int) *Int {
	val := &Int{
		&AST{
			Z3Context:	arg1.Z3Context,
			Z3AST:		C.Z3_mk_mod(arg1.Z3Context, arg1.Z3AST, arg2.Z3AST),
		},
	}

	return val;
}

// Rem returns arg1 rem arg2
// todo: finish
func (arg1 *Int) Rem(arg2 *Int) *Int {
	return nil;
}

// Power returns arg1 ^ arg2
func (arg1 *Int) Power(arg2 *Int) *Int {
	return nil;
}

// LT returns node representing t1 < t2
func (t1 *Int) LT(t2 *Int) *Int {
	val := &Int{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_lt(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return val;
}

// LE returns node representing t1 <= t2
func (t1 *Int) LE(t2 *Int) *Int {
	val := &Int{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_le(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return val;
}

// GT returns node representing t1 > t2
func (t1 *Int) GT(t2 *Int) *Int {
	val := &Int{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_gt(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return val;
}

// GE returns node representing t1 >= t2
func (t1 *Int) GE(t2 *Int) *Int {
	val := &Int{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_ge(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return val;
}

// IntToReal returns Real value from Int
func (i *Int) IntToReal() *Real {
	val := &Real{
		&AST{
			Z3Context:	i.Z3Context,
			Z3AST:		C.Z3_mk_int2real(i.Z3Context, i.Z3AST),
		},
	}

	return val;
}
