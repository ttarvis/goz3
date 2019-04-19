package main

import (
	"fmt"

	"github.com/goz3"
)

func alg1() {
	fmt.Println("*** Now running alg1 ***");
	cfg := z3.NewConfig();
	ctx := z3.NewContext(cfg);

	defer cfg.Close();
	a := ctx.MkInt(10);
	b, err := ctx.NewInt("b");

	if err != nil {
		fmt.Errorf("error: %s\n", err);
	}

	c := ctx.MkInt(20);

	ab := a.Add(b);

	// 10 + b < 20
	conjecture := ab.LT(c);

	solver := z3.NewSolver(ctx);

	solver.Assert(conjecture);

	result, err := solver.Check();
	if err != nil {
		fmt.Errorf("error: %s\n", err);
	}

	fmt.Printf("is the conjecture %s satisfiable? %t\n", conjecture, result);

	m := solver.Model();
	consts := m.GetConstInterp();

	for name, value := range consts {
		fmt.Printf("%s = %s\n", name, value);
	}

	solver.Close();
	ctx.Close();
}

func alg2() {
	fmt.Println("*** Now running alg2 ***");
        cfg := z3.NewConfig();
        ctx := z3.NewContext(cfg);

	defer cfg.Close();

	x, err := ctx.NewInt("x");
	y, err := ctx.NewInt("y");
	if err != nil {
		fmt.Errorf("error: %s", err);
	}

	z := ctx.MkInt(16);
	
	xy := x.Add(y);

	// x + y = 16
	conjecture := xy.EQ(z);

	solver := z3.NewSolver(ctx);

	solver.Assert(conjecture);

	result, err := solver.Check();
	if err != nil {
		fmt.Errorf("error: %s", err);
	}

	fmt.Println("is the conjecture (x+y = 16)satisfiable?", result);	

	m := solver.Model();
	consts := m.GetConstInterp();

        for name, value := range consts {
                fmt.Printf("%s = %s\n", name, value);
        }

        solver.Close();
        ctx.Close();
}

