package brace

import (
	"io"

	"sourcecode.social/reiver/go-utf8"
)

// ParseToBytes is similar to Parse except it writes the brace-string literal as an io.Writer.
func ParseToWriter(writer io.Writer, runescanner io.RuneScanner) error {
	if nil == writer {
		return errNilWriter
	}

	fn := func(r rune) error {
		_, err := utf8.WriteRune(writer, r)
		if nil != err {
			return err
		}

		return nil
	}

	return Parse(fn, runescanner)
}
