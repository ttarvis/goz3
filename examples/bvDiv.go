package main

import (
	"fmt"

	"github.com/goz3"
)

func bvDiv() {
	cfg := z3.NewConfig();
	ctx := z3.NewContext(cfg)
	
	defer cfg.Close();
	a, err := ctx.NewBV("a", 32);
	b, err := ctx.NewBV("b", 32);

	two := ctx.MkBV(2, 32);

	if err != nil {
		fmt.Errorf("error: %s\n", err);
	}

	uAvg := a.Add(b).Udiv(two);

	fmt.Println(uAvg);

	sAvg := a.Add(b).Sdiv(two);

	fmt.Println(sAvg);

	solver := z3.NewSolver(ctx);

	conjecture := a.Ugt(uAvg);

	fmt.Println("conjecture", conjecture);
	solver.Assert(conjecture);

	result, err := solver.Check();
	if err != nil {
		fmt.Errorf("error: %s\n", err);
	}

	fmt.Println("is the conjecture satisfiable?", result);

	m := solver.Model();
	consts := m.GetConstInterp();

	// this test should output an a and b s.t. a + b > 2^32
	for name, value := range consts {
		fmt.Printf("%s = %s\n", name, value);
	} 

	solver.Close();
	ctx.Close();
}
