package brace

import (
	"io"
)

// ParseToBytes is similar to Parse except it return the brace-string literal as a []byte.
func ParseToBytes(runescanner io.RuneScanner) ([]byte, error) {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	fn := func(r rune) error {
		p = append(p, string(r)...)
		return nil
	}

	err := Parse(fn, runescanner)
	if nil != err {
		return nil, err
	}

	return p, nil
}
