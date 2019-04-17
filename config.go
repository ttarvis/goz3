package z3


/*
#cgo LDFLAGS: -lz3
#include <z3.h>
*/
import "C"

type Config struct {
	Z3Config C.Z3_config;
}

// NewConfig makes and returns a new Config struct with a C.Z3_config object inside it.
func NewConfig() *Config {
	
	Z3Config := C.Z3_mk_config();

	config := &Config{
		Z3Config,
	}

	return config;
}

// Close closes a config object by deallocating the internal C.Z3_config object
// to implement Closer interface, must return error
func (c *Config) Close() error {
	C.Z3_del_config(c.Z3Config);
	return nil;
}

// todo: add in methods to set param value in config object
