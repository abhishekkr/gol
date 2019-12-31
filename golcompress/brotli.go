package golcompress

import (
	"gopkg.in/kothar/brotli-go.v0/dec"
	"gopkg.in/kothar/brotli-go.v0/enc"
)

func Brotli(blob []byte) (tiny []byte, err error) {
	tiny, err = enc.CompressBuffer(nil, blob, make([]byte, 0))
	return
}

func UnBrotli(tiny []byte) (blob []byte, err error) {
	blob, err = dec.DecompressBuffer(tiny, make([]byte, 0))
	return
}
