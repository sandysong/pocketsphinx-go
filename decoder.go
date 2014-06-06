package pocketsphinx

/*
#cgo pkg-config: sphinxbase pocketsphinx
#include <pocketsphinx.h>

*/
import "C"

import (
    "unsafe"
)

type Decoder C.ps_decoder_t

func NewDecoder(config *Config) *Decoder {
    return (*Decoder)(C.ps_init((*C.cmd_ln_t)(config)))
}

func (p *Decoder) ProcessRaw(data []byte, number int32, no_search int32, full_utt int32) int32 {
    searched := C.ps_process_raw((*C.ps_decoder_t)(p), (*C.int16)(unsafe.Pointer(&data[0])), C.size_t(number), C.int(no_search), C.int(full_utt))
    return int32(searched)
}

func (p *Decoder) GetHyp() (result string, score int32, uttid string) {
    var r *C.char
    var s C.int32
    var u *C.char

    r = C.ps_get_hyp((*C.ps_decoder_t)(p), &s, &u)

    result = C.GoString(r)
    score = int32(s)
    uttid = C.GoString(u)
    return
}

func (p *Decoder) StartUtt(uttid string) int {
    u := C.CString(uttid)
    defer C.free(unsafe.Pointer(u))
    rv := C.ps_start_utt((*C.ps_decoder_t)(p), u)

    return int(rv)
}

func (p *Decoder) EndUtt() int {
    rv := C.ps_end_utt((*C.ps_decoder_t)(p))

    return int(rv)
}

func (p *Decoder) Free() int {
    res := C.ps_free((*C.ps_decoder_t)(p))

    return int(res)
}

func PsArgs() (arg *Arg) {
    a := C.ps_args()
    return (*Arg)(a)
}

