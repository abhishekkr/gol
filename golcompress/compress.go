package golcompress

import "errors"

func Compress(blob []byte, scheme string) ([]byte, error) {
	switch scheme {
	case "brotli":
		return Brotli(blob)
	default:
		return []byte{}, errors.New("unsupported compression scheme desired")
	}
}

func Decompress(tiny []byte, scheme string) ([]byte, error) {
	switch scheme {
	case "brotli":
		return UnBrotli(tiny)
	default:
		return []byte{}, errors.New("unsupported compression scheme desired")
	}
}
