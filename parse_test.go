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
		{t: SymbolIdent, ident: "foo"},
		{t: SymbolBinaryOr},
		{t: SymbolIdent, ident: "bar"},
		{t: SymbolBinaryAnd},
		{t: SymbolUnaryNot},
		{t: SymbolLParen},
		{t: SymbolIdent, ident: "baz"},
		{t: SymbolBinaryAnd},
		{t: SymbolIdent, ident: "bar"},
		{t: SymbolBinaryOr},
		{t: SymbolIdent, ident: "quux"},
		{t: SymbolRParen},
	}
	assert.Equal(t, expect, tkns)
}
