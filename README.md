# go-brace

Package **brace** implements tools for working with brace-string literals, for the Go programming language.

Brace-String literals are a type of string literal used in the Tcl programming language.
For example:
```
{Hello world!}
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-brace

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-brace?status.svg)](https://godoc.org/sourcecode.social/reiver/go-brace)

## Example

Here is an example:
```golang

import "sourcecode.social/reiver/go-brace"
import "sourcecode.social/reiver/go-utf-8"

// ...

const value string = `{This is a brace-string literal. It begins with a \{ symbol and end with a \} symbol. Now look at this: {wow} }`

var reader io.Reader = strings.NewReader(value)
var runereader io.RuneReader = utf8.NewRuneReader(reader)

bracestring, err := brace.ParseToString(runereader)
if nil != err {
		return err
}

fmt.Print(bracestring)

// This will print:
//
//	This is a brace-string literal. It begins with a { symbol and end with a } symbol. Now look at this: {wow} 
//
// I.e., the string:
//
//	"This is a brace-string literal. It begins with a { symbol and end with a } symbol. Now look at this: {wow} "
//
// Notice that the beginning '{' is not there, that the ending '}' is not there, and the 2 '\' are not there either.
```

## Import

To import package **brace** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-brace"
```

## Installation

To install package **brace** do the following:
```
GOPROXY=direct go get https://sourcecode.social/reiver/go-brace
```

## Author

Package **brace** was written by [Charles Iliya Krempeaux](http://changelog.ca)
