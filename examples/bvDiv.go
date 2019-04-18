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
}
