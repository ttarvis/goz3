package z3


/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"

//export goZ3ErrorHandler
func goZ3ErrorHandler(ctx C.Z3_context, e C.Z3_error_code) {
	msg := C.Z3_get_error_msg_ex(ctx, e);
	// todo: consider not using panic and instead lift the error in
	// to a good Go error.
	panic(C.GoString(msg));
}
