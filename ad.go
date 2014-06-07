package pocketsphinx

/*
#cgo pkg-config: sphinxbase
#include <sphinxbase/ad.h>
#include <sphinxbase/cont_ad.h>

typedef int32 (*adfunc)(ad_rec_t *ad, int16 *buf, int32 max);

int32 cont_read_ts(cont_ad_t *cont) {
    return cont->read_ts;
}

*/
import "C"

import (
    "errors"
    "unsafe"
)

type Rec C.ad_rec_t

func OpenRec() *Rec {
    return (*Rec)(C.ad_open())

}

func OpenRecSps(samples int32) *Rec {
    return (*Rec)(C.ad_open_sps(C.int32(samples)))
}

func OpenRecDev(dev string, samples int32) *Rec {
    d := C.CString(dev)
    defer C.free(unsafe.Pointer(d))

    return (*Rec)(C.ad_open_dev(d, C.int32(samples)))
}


func (rec *Rec) StartRec() error {
    err := C.ad_start_rec((*C.ad_rec_t)(rec))
    if err != 0 {
        return errors.New("Cannot start record.")
    }
    return nil
}

func (rec *Rec) StopRec() error {
    err := C.ad_stop_rec((*C.ad_rec_t)(rec))
    if err != 0 {
        return errors.New("Cannot stop record.")
    }
    return nil
}

func (rec *Rec) Close() error {
    err := C.ad_close((*C.ad_rec_t)(rec))
    if err != 0 {
        return errors.New("Cannot close record.")
    }
    return nil
}

func (rec *Rec) Read(buf []int16, max int32) int32 {
    return int32(C.ad_read((*C.ad_rec_t)(rec), (*C.int16)(unsafe.Pointer(&buf[0])), C.int32(max)))
}

type ContAd C.cont_ad_t

func InitCont(rec *Rec) *ContAd {
    return (*ContAd)(C.cont_ad_init((*C.ad_rec_t)(rec), C.adfunc(C.ad_read)))
}

func (cont *ContAd) ContCalibrate() error {
    err := C.cont_ad_calib((*C.cont_ad_t)(cont))
    if err != 0 {
        return errors.New("Cannot calibrate cont_ad.")
    }
    return nil
}

func (cont *ContAd) ContRead(buf []int16, max int32) int32 {
    return int32(C.cont_ad_read((*C.cont_ad_t)(cont), (*C.int16)(unsafe.Pointer(&buf[0])), C.int32(max)))
}

func (cont *ContAd) ContReadTs() int32 {
    return int32(C.cont_read_ts((*C.cont_ad_t)(cont)))
}

func (cont *ContAd) ContClose() error {
    err := C.cont_ad_close((*C.cont_ad_t)(cont))
    if err != 0 {
        return errors.New("Cannot close cont_ad.")
    }
    return nil
}

func (cont *ContAd) Reset() int32 {
    return int32(C.cont_ad_reset((*C.cont_ad_t)(cont)))
}
