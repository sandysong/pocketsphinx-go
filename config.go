package pocketsphinx

/*
#cgo pkg-config: sphinxbase
#include <stdlib.h>
#include <sphinxbase/cmd_ln.h>

*/
import "C"

import(
    "unsafe"
)

type Arg C.arg_t

type Config C.cmd_ln_t

func NewConfig(config *Config, arg *Arg, strict int32) *Config {
    return (*Config)(C.cmd_ln_parse_r((*C.cmd_ln_t)(config), (*C.arg_t)(arg), C.int32(0), nil, C.int32(strict)))
}

func (c *Config) SetInt(key string, value int64) {
    k := C.CString(key)
    defer C.free(unsafe.Pointer(k))

    C.cmd_ln_set_int_r((*C.cmd_ln_t)(c), k, C.long(value))
}

func (c *Config) SetStr(key string, value string) {
    k := C.CString(key)
    defer C.free(unsafe.Pointer(k))
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))

    C.cmd_ln_set_str_r((*C.cmd_ln_t)(c), k, v)
}

func (c *Config) SetFloat(key string, value float64) {
    k := C.CString(key)
    defer C.free(unsafe.Pointer(k))

    C.cmd_ln_set_float_r((*C.cmd_ln_t)(c), k, C.double(value))
}

func (c *Config) Free() int32 {
    res := C.cmd_ln_free_r((*C.cmd_ln_t)(c))
    c = nil
    return int32(res)
}
