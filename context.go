package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>

extern void goZ3ErrorHandler(Z3_context c, Z3_error_code e);
*/
import "C"

// todo: figure out what a context really is
type Context struct {
	Z3Context C.Z3_context;
}

func NewContext(c *Config) *Context {
	ctx := &Context{
		Z3Context: C.Z3_mk_context(c.Z3Config),
	}

	// todo: you should set a finalizer to deallocate the context just in case
	
	// install a Z3 error handler
	// convert the Go error handler to a C error handler
	// this can be done because of the extern statement in error.go
	C.Z3_set_error_handler(ctx.Z3Context,(*C.Z3_error_handler)(C.goZ3ErrorHandler));
	return ctx
}

func (ctx *Context) Close() error {
	// Clear context
	C.Z3_del_context(ctx.Z3Context);

	// todo: if the Context is closed, what to do with error handlers set up for it?
	return nil;
}

/*
func (ctx *Context) Interrupt() {
	C.Z3_interrupt(ctx.Z3Config);
	todo: runtime.KeepAlive(ctx);
}
*/	
