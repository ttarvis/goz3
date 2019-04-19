package main

import (
	"fmt"

	"github.com/goz3"
)

func divBug() {
	cfg := z3.NewConfig();
	ctx := z3.NewContext(cfg)
	
	defer cfg.Close();
	a, err := ctx.NewBV("a", 32);
	b, err := ctx.NewBV("b", 32);
	if err != nil {
		fmt.Errorf("error: %s", err);
	}

	two := ctx.MkBV(2, 32);

	// (a + b)
	ab := a.Add(b);
	// (a + b) / 2
	uDiv := ab.Udiv(two); 
	uDiv.Add(ab);

	//uDiv2 := ((a.Zeroext(1)).Add(b.Zeroext(1)).Udiv(two)).Extract(31,0);
	uDiv2 := ((a.Zeroext(1)).Add(b.Zeroext(1)).Udiv(two.Zeroext(1)))
	ext := uDiv2.Extract(31,0);

	conjecture := uDiv.NE(ext);
	fmt.Println(conjecture);

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
