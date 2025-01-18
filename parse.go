package brace

import (
	"io"

	"github.com/reiver/go-erorr"
)

// Parse parses a brace-string literal from a io.RuneScanner.
//
// Parse will call ‘fn’ for each logical character it receives.
// So — if this is the brace-literal it parses:
//
//	`{a b \{ c \} d}`
//
// Then ‘fn’ would be called 11 times, and given these runes:
//
//	'a'
//	' '
//	'b'
//	' '
//	'{'
//	' '
//	'c'
//	' '
//	'}'
//	' '
//	'd'
//
// Note that the beginning '{' and ending '}' of the brace-string literal are not part of this.
// Also note that the '\' (black-slash) before the '{' and the '}' was not included either.
func Parse(fn func(rune)error, runescanner io.RuneScanner) error {

	if nil == runescanner {
		return errNilRuneScanner
	}

	var nextRead uint64 = 1

	{
		r, _, err := runescanner.ReadRune()
		if io.EOF == err {
			return errProblemReadingCharacterNumber(nextRead, errEmptyString)
		}
		if nil != err {
			return errProblemReadingCharacterNumber(nextRead, err)
		}
		nextRead++

		{
			const expected = leftbrace
			actual := r

			if expected != actual {
				err := runescanner.UnreadRune()
				if nil != err {
					characterNumber := nextRead - 1
					return errProblemUnreadingCharacterNumber(characterNumber, err, r)
				}

				return erorr.Errorf("brace: expected first character of brace-string literal to be a %q (%U), but actually was %q (%U)", expected, expected, actual, actual)
			}
		}
	}

	var depth int64 = 1
	loop: for {
		r, _, err := runescanner.ReadRune()
		if io.EOF == err {
			return errProblemReadingCharacterNumber(nextRead, errFileEndedBeforeBraceStringLiteralClosed(depth))
		}
		if nil != err {
			return errProblemReadingCharacterNumber(nextRead, err)
		}
		nextRead++

		switch r {
		case leftbrace:
			depth++
		case rightbrace:
			depth--
		case '\\':
			r, _, err = runescanner.ReadRune()
			if io.EOF == err {
				return errProblemReadingCharacterNumber(nextRead, errFileEndedBeforeBraceStringLiteralClosed(depth))
			}
			if nil != err {
				return erorr.Errorf("brace: problem reading character №%d (which would have followed a black-slash %q (%U) that was just read) of what should have been a brace-string literal: %w", nextRead, r, r, err)
			}
			nextRead++
		}

		switch {
		case 0 == depth:
	////////////// BREAK
			break loop
		case depth < 0:
			return errInternalError(errParserDepthNegative(depth))
		}

		{
			err := fn(r)
			if nil != err {
				return err
			}
		}
	}

	return nil
}
