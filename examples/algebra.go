package main

import (
	"fmt"

	"github.com/goz3"
)

func alg1() {
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

	conjecture := ab.LT(c);

	solver := z3.NewSolver(ctx);

	solver.Assert(conjecture);

	result, err := solver.Check();
	if err != nil {
		fmt.Errorf("error: %s\n", err);
	}

	fmt.Println("is the conjecture satisfiable?", result);

	m := solver.Model();
	consts := m.GetConstInterp();

	for name, value := range consts {
		fmt.Printf("%s = %s\n", name, value);
	}

	solver.Close();
	ctx.Close();
}
