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

type Solver struct {
	Z3Context	C.Z3_context;
	Z3Solver	C.Z3_solver;
}

// NewSolver makes and returns a new solver.
func NewSolver(ctx *Context) *Solver {
	solver := &Solver{
		Z3Context:	ctx.Z3Context,
		Z3Solver:	C.Z3_mk_solver(ctx.Z3Context),
	}

	C.Z3_solver_inc_ref(solver.Z3Context, solver.Z3Solver);

	return solver;
}

// Assert uses Z3 to assert a conjecture, conj.
// It returns nothing.  Use Check to check the satisfiability
func (s *Solver) Assert(conj value) {
	C.Z3_solver_assert(s.Z3Context, s.Z3Solver, conj.GetAST());
}

// Check determines whether the conjecture in solver s is satisfiable.
func (s *Solver) Check() (bool, error) {
	// Z3 lifted bool: true, undef, false
	var res C.Z3_lbool;
	var err error;
	res = C.Z3_solver_check(s.Z3Context, s.Z3Solver);
	if res == C.Z3_L_UNDEF {
		Z3error := C.Z3_solver_get_reason_unknown(s.Z3Context, s.Z3Solver);
		err = errors.New(C.GoString(Z3error));
	}
	// note: cannot use switch statement here with C types
	/*
	other cases
	(these are lifted booleans)
	C.Z3_L_FALSE
	C.Z3_L_TRUE
	*/

	return res == C.Z3_TRUE, err;
}

// Close frees a solver's memory
// it does this by decrementing the solver's reference
// implement Closer interface
// todo: what does that mean
func (s *Solver) Close() error {
	C.Z3_solver_dec_ref(s.Z3Context, s.Z3Solver);
	return nil;
}

// Model returns the last model from Check
func (s *Solver) Model() *Model {
	model := &Model{
		Z3Context:	s.Z3Context,
		Z3Model:	C.Z3_solver_get_model(s.Z3Context, s.Z3Solver),
	}

	return model;
}
