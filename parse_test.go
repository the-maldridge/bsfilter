package bsfilter

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	// The spacing is deliberately changed in this to ensure that
	// it does not get bothered by whitespace.
	tkns := Tokenize("foo | bar&!( baz&bar|quux  )")

	expect := []Symbol{
		{T: SymbolIdent, Ident: "foo"},
		{T: SymbolBinaryOr},
		{T: SymbolIdent, Ident: "bar"},
		{T: SymbolBinaryAnd},
		{T: SymbolUnaryNot},
		{T: SymbolLParen},
		{T: SymbolIdent, Ident: "baz"},
		{T: SymbolBinaryAnd},
		{T: SymbolIdent, Ident: "bar"},
		{T: SymbolBinaryOr},
		{T: SymbolIdent, Ident: "quux"},
		{T: SymbolRParen},
	}
	assert.Equal(t, expect, tkns)
}

func TestNewInvalid(t *testing.T) {
	e, err := New("")
	assert.Nil(t, e)
	assert.NotNil(t, err)
}

func TestNew(t *testing.T) {
	e, err := New("foo&bar&!(baz|quux)")
	assert.Nil(t, err)

	expected := &Expression{
		l: log.Default(),
		root: &ASTNode{
			Symbol: &Symbol{T: SymbolBinaryAnd},
			Left:   &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "foo"}},
			Right: &ASTNode{
				Symbol: &Symbol{T: SymbolBinaryAnd},
				Left:   &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "bar"}},
				Right: &ASTNode{
					Symbol: &Symbol{T: SymbolUnaryNot},
					Right: &ASTNode{
						Symbol: &Symbol{T: SymbolBinaryOr},
						Left:   &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "baz"}},
						Right:  &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "quux"}},
					},
				},
			},
		},
	}
	assert.Equal(t, expected, e)
}
