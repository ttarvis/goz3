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

// BV wraps an AST node to make a Bit Vector type
type BV struct {
	*AST
}

// NewBV makes and returns a Bit Vector of size sz or an error
func (ctx *Context) NewBV(name interface{}, sz uint) (*BV, error) {
	bvSort := ctx.MakeBV(sz);
	var varSym *Symbol;
	switch v := name.(type) {
	case int:
		varSym = ctx.IntSymbol(v);
	case string:
		varSym = ctx.StringSymbol(v);
	default:
		return nil, errors.New("Int Const name type must be int or string");
	}

	bvConst := &BV{
		ctx.Const(varSym, bvSort),
	}

	return bvConst, nil;
}

// Not is a bitwise negation of a bit vector
func (bv *BV) Not() *BV {
	v := &BV{
		&AST{
			Z3Context:	bv.Z3Context,
			Z3AST:		C.Z3_mk_bvnot(bv.Z3Context, bv.Z3AST),
		},
	}

	return v;
}

// Redand takes a bit vector and returns a vector of conjunction of bits 
func (bv *BV) Redand() *BV {
	v := &BV{
		&AST{
			Z3Context:	bv.Z3Context,
			Z3AST:		C.Z3_mk_bvredand(bv.Z3Context, bv.Z3AST),
		},
	}

	return v;
}

// Redor takes disjunction of bits in vector, return vector of length 1 
func (bv *BV) Redor() *BV {
	v := &BV{
		&AST{
			Z3Context:	bv.Z3Context,
			Z3AST:		C.Z3_mk_bvredor(bv.Z3Context, bv.Z3AST),
		},
	}

	return v;
}

// And does bitwise and on t1, t2. Returns vector = t1 AND t2
func (t1 *BV) And(t2 *BV) *BV {
	v := &BV{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_bvand(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return v;
}

// Or does bitwise or.  Returns vector = t1 OR t2
func (t1 *BV) Or(t2 *BV) *BV {
	v := &BV{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_bvor(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return v;
}

// Xor returns bitwise exclusive-or. t1 and t2 must have same length
func (t1 *BV) Xor(t2 *BV) *BV {
	v := &BV{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_bvxor(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return v;
}

// Nand returns bitwise nand.  t1 and t2 must have same length
func (t1 *BV) Nand(t2 *BV) *BV {
	v := &BV{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_bvnand(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return v;
}

// Nor returns bitwise nor.
func (t1 *BV) Nor(t2 *BV) *BV {
	v := &BV{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_bvnor(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return v;
}

// Xnor returns bitwise xnor
func (t1 *BV) Xnor(t2 *BV) *BV {
	v := &BV{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_bvxnor(t1.Z3Context, t1.Z3AST, t2.Z3AST),
		},
	}

	return v;
}

// Neg is two's complement unary negative operation.  Neg returns -(bv) 
func (bv *BV) Neg() *BV {
        v := &BV{
                &AST{
                        Z3Context:      bv.Z3Context,
                        Z3AST:          C.Z3_mk_bvneg(bv.Z3Context, bv.Z3AST),
                },
        }

        return v;
}

// Add does two's complement addition.  It adds t1 to t2 and returns a bit vector
func (t1 *BV) Add(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvadd(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Sub does standard two's complement subtraction.  t1 - t2.
func (t1 *BV) Sub(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsub(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Mul does standard two's complement multiplication
func (t1 *BV) Mul(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvmul(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Udiv does unsigned division
func (t1 *BV) Udiv(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvudiv(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Sdiv does two's complement signed division
func (t1 *BV) Sdiv(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsdiv(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Urem returns unsigned remainder
func (t1 *BV) Urem(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvurem(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Srem does two's complement signed division.
// sign follows dividend
func (t1 *BV) Srem(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsrem(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Smod does two's complement signed remainder (sign follows divisor)
func (t1 *BV) Smod(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsmod(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Ult is unsigned less than t1 < t2
func (t1 *BV) Ult(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvult(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Ule is unsigned less than or equal to t1 <= t2
func (t1 *BV) Ule(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvule(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Sle is two's complement signed less than or equal to
func (t1 *BV) Sle(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsle(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Uge is unsigned greater than or equal to
func (t1 *BV) Uge(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvuge(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Sge is two's complement signed greater than or equal to
func (t1 *BV) Sge(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsge(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Ugt is unsigned greater than
func (t1 *BV) Ugt(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvugt(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Sgt is two's complement signed greater than or equal to
func (t1 *BV) Sgt(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsgt(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Concat concatenates the given bit vectors
func (t1 *BV) Concat(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_concat(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Extract extracts the bits from high down to low from a bit vector of size m to
// yield a new bit-vector of size n. where n = high - low + 1.
func (bv *BV) Extract(high, low uint) *BV {
	v := &BV{
                &AST{
                        Z3Context:      bv.Z3Context,
                        Z3AST:          C.Z3_mk_extract(bv.Z3Context, C.uint(high), C.uint(low), bv.Z3AST),
                },
        }

	return v;
}

// Signext Sign-extend of the given bit-vector to the (signed) equivalent bit-vector of size m+i,
// where m is the size of the given bit-vector.
func (bv *BV) Signext(i uint) *BV {
	v := &BV{
		&AST{
			Z3Context:	bv.Z3Context,
			Z3AST:		C.Z3_mk_sign_ext(bv.Z3Context, C.uint(i), bv.Z3AST),
		},
	}

	return v;
}

// Zeroext extends given bit vector with zeros to the (unsigned) equivalent bit-vector of size m+i,
// where m is the size of the given bit-vector.
func (bv *BV) Zeroext(i uint) *BV {
        v := &BV{
                &AST{
                        Z3Context:      bv.Z3Context,
                        Z3AST:          C.Z3_mk_zero_ext(bv.Z3Context, C.uint(i), bv.Z3AST),
                },
        }

        return v;
}

// Repeat repeats a given bit vector up length i
func (bv *BV) Repeat(i uint) *BV {
        v := &BV{
                &AST{
                        Z3Context:      bv.Z3Context,
                        Z3AST:          C.Z3_mk_repeat(bv.Z3Context, C.uint(i), bv.Z3AST),
                },
        }

        return v;
}

// Shl shifts left.  It is equivalent to multiplication by 2^x where x is the value of t2.
func (t1 *BV) Shl(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvshl(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Lshr does logical shift right.  It is equivalent to unsigned division by 2^x where x is t2
func (t1 *BV) Lshr(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvlshr(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// Ashr is arithmetic shift right.  
func (t1 *BV) Ashr(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvashr(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// RotateLeft rotates bits of bv to the left i times.
func (bv *BV) RotateLeft(i uint) *BV {
        v := &BV{
                &AST{
                        Z3Context:      bv.Z3Context,
                        Z3AST:          C.Z3_mk_rotate_left(bv.Z3Context, C.uint(i), bv.Z3AST),
                },
        }

        return v;
}

// RotateRight
func (bv *BV) RotateRight(i uint) *BV {
        v := &BV{
                &AST{
                        Z3Context:      bv.Z3Context,
                        Z3AST:          C.Z3_mk_rotate_right(bv.Z3Context, C.uint(i), bv.Z3AST),
                },
        }

        return v;
}

// ExtRotateLeft rotates bits of t1 to the left t2 times
func (t1 *BV) ExtRotateLeft(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_ext_rotate_left(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// ExtRotateRight rotates bits of t1 to the right t2 times
func (t1 *BV) ExtRotateRight(t2 *BV) *BV {
        v := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_ext_rotate_right(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return v;
}

// BVToInt returns an Int from a bit vector.  If isSigned is false, then the bit vector bv
// is treated as unsigned.  So the result is non-negative and in the rage [0..2^N-1], where
// N is the number of bits in bv.  If isSigned is true, t1 is treated as a signed bit vector
func (bv *BV) BVToInt(signed bool) *Int {
	var isSigned C.Z3_bool
	if signed {
		isSigned = C.Z3_TRUE;
	} else {
		isSigned = C.Z3_FALSE;
	}
	val := &Int{
		&AST{
			Z3Context:	bv.Z3Context,
			Z3AST:		C.Z3_mk_bv2int(bv.Z3Context, bv.Z3AST, isSigned),
		},
	}

	return val;
}

// AddNoOverflow checks that t1 + t2 does not overflow
func (t1 * BV) AddNoOverflow(t2 *BV, signed bool) *BV {
	var isSigned C.Z3_bool;
	if signed {
		isSigned = C.Z3_TRUE;
	} else {
		isSigned = C.Z3_FALSE;
	}
	val := &BV{
		&AST{
			Z3Context:	t1.Z3Context,
			Z3AST:		C.Z3_mk_bvadd_no_overflow(t1.Z3Context, t1.Z3AST, t2.Z3AST, isSigned),
		},
	}

	return val;
}

// AddNoUnderflow checks that bit-wise signed addition of t1 and t2 does not underflow
func (t1 * BV) AddNoUnderflow(t2 *BV) *BV {
        val := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvadd_no_underflow(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return val;
}

// SubNoOverflow checks that the bit wise signed subtraction of t1 and t2 does not overflow
func (t1 * BV) SubNoOverflow(t2 *BV) *BV {
        val := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsub_no_overflow(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return val;
}

func (t1 * BV) SubNoUnderflow(t2 *BV, signed bool) *BV {
        var isSigned C.Z3_bool;
        if signed {
                isSigned = C.Z3_TRUE;
        } else {
                isSigned = C.Z3_FALSE;
        }
        val := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsub_no_underflow(t1.Z3Context, t1.Z3AST, t2.Z3AST, isSigned),
                },
        }

        return val;
}

// SdivNoOverflow checks that the bit wise signed division of t1 and t2 does not overflow
func (t1 * BV) SdivNoOverflow(t2 *BV) *BV {
        val := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvsdiv_no_overflow(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return val;
}

// NegNoOverflow checks that bit wise negation does not overflow when t1 is interpreted as a signed bit vector
func (t1 * BV) NegNoOverflow() *BV {
        val := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvneg_no_overflow(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return val;
}

// MulNoOverflow checks that the bit wise multiplication of t1 and t2 does not overflow
func (t1 * BV) MulNoOverflow(t2 *BV, signed bool) *BV {
        var isSigned C.Z3_bool;
        if signed {
                isSigned = C.Z3_TRUE;
        } else {
                isSigned = C.Z3_FALSE;
        }
        val := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvmul_no_overflow(t1.Z3Context, t1.Z3AST, t2.Z3AST, isSigned),
                },
        }

        return val;
}

// MulNoUnderflow checks that bit wise signed multiplication of t1 and t2 does not underflow
func (t1 * BV) MulNoUnderflow(t2 *BV) *BV {
        val := &BV{
                &AST{
                        Z3Context:      t1.Z3Context,
                        Z3AST:          C.Z3_mk_bvmul_no_underflow(t1.Z3Context, t1.Z3AST, t2.Z3AST),
                },
        }

        return val;
}
