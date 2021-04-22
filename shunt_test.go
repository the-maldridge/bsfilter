package bsfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	// The spacing is deliberately changed in this to ensure that
	// it does not get bothered by whitespace.
	tkns := Tokenize("foo | bar&!( baz&bar|quux  )")

	expect := []Symbol{
		{t: symbolIdent, ident: "foo"},
		{t: symbolBinaryOr},
		{t: symbolIdent, ident: "bar"},
		{t: symbolBinaryAnd},
		{t: symbolUnaryNot},
		{t: symbolLParen},
		{t: symbolIdent, ident: "baz"},
		{t: symbolBinaryAnd},
		{t: symbolIdent, ident: "bar"},
		{t: symbolBinaryOr},
		{t: symbolIdent, ident: "quux"},
		{t: symbolRParen},
	}
	assert.Equal(t, expect, tkns)
}
