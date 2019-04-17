package z3

import "testing"

// TestDeMorgan tests the negation of De Morgan's Laws
func TestDeMorgan(t *testing.T) {
	cfg	:= NewConfig();
	ctx	:= NewContext(cfg);
	solver	:= NewSolver(ctx);
	bsort	:= ctx.MakeBool();
	symX	:= ctx.StringSymbol("x");
	symY	:= ctx.StringSymbol("y");
	x	:= ctx.Const(symX, bsort);
	y	:= ctx.Const(symY, bsort);

	defer cfg.Close();

	t.Logf("De Morgan's Law test\n");

	/* De Morgan - with a negation around */
	/* !(!(x && y) <-> (!x || !y)) */
	ls	:= x.And(y).Not();
	rs	:= x.Not().Or(y.Not());

	t.Logf("ls is %s", ls.String());
	t.Logf("rs is %s", rs.String());

	conjecture		:= ls.IFF(rs);
	negatedConjecture	:= conjecture.Not();

	solver.Assert(negatedConjecture);
	sat, err := solver.Check();
	if err != nil {
		t.Errorf("negated conjecture failed is satisfiable: %s", err);
	} else if sat {
		t.Errorf("negation of De Morgan's law is satisfiable, hence disproved De Morgan's Law");
	}

	solver.Close();	
	ctx.Close();
}

func TestDeMorgan2(t *testing.T) {
	cfg	:= NewConfig();
	ctx	:= NewContext(cfg);	
	solver	:= NewSolver(ctx);

	defer cfg.Close();

	x, err := ctx.NewBool("x");
	y, err := ctx.NewBool("y");
	if err != nil {
		t.Errorf("failed to properly allocate x and y");
	}

	t.Logf("De Morgan's Law Test 2\n");

        /* De Morgan - with a negation around */
        /* !(!(x && y) <-> (!x || !y)) */
	ls := x.And(y).Not();
	rs := x.Not().Or(y.Not());

	t.Logf("ls is %s", ls.String());
	t.Logf("rs is %s", rs.String());

	conjecture		:= ls.Iff(rs);
	negatedConjecture	:= conjecture.Not();

	solver.Assert(negatedConjecture);
	sat, err := solver.Check();
	if err != nil {
		t.Errorf("negated conjecture failed is satisfiable: %s", err);
	} else if sat {
		t.Errorf("negation of De Morgan's law is satisfiable, hence disproved De Morgan's Law");
	}

	solver.Close();
	ctx.Close();
}
