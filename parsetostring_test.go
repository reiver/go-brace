package brace_test

import (
	"testing"

	"io"
	"strings"

	"sourcecode.social/reiver/go-utf8"

	"sourcecode.social/reiver/go-brace"
)

func TestParseToString(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
			Value:   `{}`,
			Expected: ``,
		},



		{
			Value:   `{apple}`,
			Expected: `apple`,
		},
		{
			Value:   `{banana}`,
			Expected: `banana`,
		},
		{
			Value:   `{cherry}`,
			Expected: `cherry`,
		},



		{
			Value:   `{\{}`,
			Expected:  `{`,
		},
		{
			Value:   `{\}}`,
			Expected:  `}`,
		},



		{
			Value:   `{$x}`,
			Expected: `$x`,
		},
		{
			Value:   `{$xy}`,
			Expected: `$xy`,
		},
		{
			Value:   `{$xyz}`,
			Expected: `$xyz`,
		},



		{
			Value:   `{ set var 123 }`,
			Expected: ` set var 123 `,
		},
		{
			Value:   `{ set var {123} }`,
			Expected: ` set var {123} `,
		},
		{
			Value:   `{ set {var} {123} }`,
			Expected: ` set {var} {123} `,
		},
		{
			Value:   `{ set var {{1}{2}{3}} }`,
			Expected: ` set var {{1}{2}{3}} `,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		expected := test.Expected
		actual, err := brace.ParseToString(runescanner)
		if nil != err  {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		if expected != actual {
			t.Errorf("For test #%d, the actual parsed value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}

func TestParseToString_fail(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedErr string
	}{
		{
			Value: ``,
			ExpectedErr: `brace: problem reading character №1 of what should have been a brace-string literal: empty string`,
		},



		{
			Value:                                                                                                       `apple`,
			ExpectedErr: `brace: expected first character of brace-string literal to be a '{' (U+007B), but actually was 'a' (U+0061)`,
		},
		{
			Value:                                                                                                       `banana`,
			ExpectedErr: `brace: expected first character of brace-string literal to be a '{' (U+007B), but actually was 'b' (U+0062)`,
		},
		{
			Value:                                                                                                       `cherry`,
			ExpectedErr: `brace: expected first character of brace-string literal to be a '{' (U+007B), but actually was 'c' (U+0063)`,
		},



		{
			Value: `{`,
			ExpectedErr: `brace: problem reading character №2 of what should have been a brace-string literal: file ended before the brace-string literal was closed — expected 1 more '}' (U+007D) character(s)`,
		},
		{
			Value:                                                                                                       `}`,
			ExpectedErr: `brace: expected first character of brace-string literal to be a '{' (U+007B), but actually was '}' (U+007D)`,
		},



		{
			Value:                                                                                                       `"`,
			ExpectedErr: `brace: expected first character of brace-string literal to be a '{' (U+007B), but actually was '"' (U+0022)`,
		},



		{
			Value: "{ this is {a",
			ExpectedErr: "brace: problem reading character №13 of what should have been a brace-string literal: file ended before the brace-string literal was closed — expected 2 more '}' (U+007D) character(s)",
		},
		{
			Value: "{ this {is} {a",
			ExpectedErr: "brace: problem reading character №15 of what should have been a brace-string literal: file ended before the brace-string literal was closed — expected 2 more '}' (U+007D) character(s)",
		},
		{
			Value: "{ {thi}s i{s} {a",
			ExpectedErr: "brace: problem reading character №17 of what should have been a brace-string literal: file ended before the brace-string literal was closed — expected 2 more '}' (U+007D) character(s)",
		},
		{
			Value: "{ this is {a {dangling",
			ExpectedErr: "brace: problem reading character №23 of what should have been a brace-string literal: file ended before the brace-string literal was closed — expected 3 more '}' (U+007D) character(s)",
		},
		{
			Value: "{ this {is} {a {dangling",
			ExpectedErr: "brace: problem reading character №25 of what should have been a brace-string literal: file ended before the brace-string literal was closed — expected 3 more '}' (U+007D) character(s)",
		},
		{
			Value: "{ {thi}s i{s} {a {dangling",
			ExpectedErr: "brace: problem reading character №27 of what should have been a brace-string literal: file ended before the brace-string literal was closed — expected 3 more '}' (U+007D) character(s)",
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)
		var runescanner io.RuneScanner = utf8.NewRuneScanner(reader)

		_, err := brace.ParseToString(runescanner)
		if nil == err  {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("ERROR:     (%T) %s", err, err)
			t.Logf("EXPECTED-ERROR: %q", test.ExpectedErr)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			expected := test.ExpectedErr
			actual := err.Error()

			if expected != actual {
				t.Errorf("For test #%d, the actual error value is not what was expected.", testNumber)
				t.Logf("EXPECTED-ERROR: %q", expected)
				t.Logf("ACTUAL-ERROR:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
