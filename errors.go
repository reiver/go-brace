package brace

import (
	"github.com/reiver/go-erorr"
)

const (
	errEmptyString    = erorr.Error("empty string")
	errNilRuneScanner = erorr.Error("brace: nil rune-scanner")
	errNilWriter      = erorr.Error("brace: nil writer")
)

func errFileEndedBeforeBraceStringLiteralClosed(depth int64) error {
	return erorr.Errorf("file ended before the brace-string literal was closed — expected %d more %q (%U) character(s)", depth, rightbrace, rightbrace)
}

func errInternalError(err error) error {
	return erorr.Errorf("brace: internal error: %w", err)
}

func errParserDepthNegative(depth int64) error {
	return erorr.Errorf("brace: parser depth (%d) is negative", depth)
}

func errProblemReadingCharacterNumber(num uint64, err error) error {
	return erorr.Errorf("brace: problem reading character №%d of what should have been a brace-string literal: %w", num, err)
}

func errProblemUnreadingCharacterNumber(num uint64, err error, r rune) error {
	return erorr.Errorf("brace: problem unreading character №%d (which is a %q (%U)) of what should have been a brace-string literal: %w", num, r, r, err)
}
