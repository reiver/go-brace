package brace

import (
	"io"
)

// ParseToString is similar to Parse except it return the brace-string literal as a string.
func ParseToString(runescanner io.RuneScanner) (string, error) {
	p, err := ParseToBytes(runescanner)
	if nil != err {
		return "", err
	}

	return string(p), nil
}
