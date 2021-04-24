package bsfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestASTResolve(t *testing.T) {
	cases := []struct {
		node ASTNode
		want bool
	}{
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolUnaryNot},
				Right:  &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
			},
			want: false,
		},
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"},
			},
			want: true,
		},
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolBinaryOr},
				Right:  &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testFalse"}},
				Left:   &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
			},
			want: true,
		},
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolBinaryAnd},
				Right:  &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
				Left:   &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
			},
			want: true,
		},
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolLParen},
			},
			want: false,
		},
	}
	vset := ValueSet{"testTrue": struct{}{}}

	for i, c := range cases {
		res := c.node.Resolve(vset)
		assert.Equalf(t, c.want, res, "Case %d: Got %v want %v", i, c.want, res)
	}
}

func TestASTString(t *testing.T) {
	cases := []struct {
		node ASTNode
		want string
	}{
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolUnaryNot},
				Right:  &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
			},
			want: "!testTrue",
		},
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"},
			},
			want: "testTrue",
		},
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolBinaryOr},
				Right:  &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testFalse"}},
				Left:   &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
			},
			want: "(testTrue|testFalse)",
		},
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolBinaryAnd},
				Right:  &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
				Left:   &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
			},
			want: "(testTrue&testTrue)",
		},
		{
			node: ASTNode{
				Symbol: &Symbol{T: SymbolLParen},
			},
			want: "",
		},
	}

	for i, c := range cases {
		res := c.node.String()
		assert.Equalf(t, c.want, res, "Case %d: Got %v want %v", i, c.want, res)
	}
}

func TestExpressionString(t *testing.T) {
	node := ASTNode{
		Symbol: &Symbol{T: SymbolBinaryAnd},
		Right:  &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
		Left:   &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
	}
	want := "(testTrue&testTrue)"

	assert.Equal(t, want, Expression{root: &node}.String())

}

func TestExpressionEvaluate(t *testing.T) {
	node := ASTNode{
		Symbol: &Symbol{T: SymbolBinaryAnd},
		Right:  &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testTrue"}},
		Left:   &ASTNode{Symbol: &Symbol{T: SymbolIdent, Ident: "testFalse"}},
	}
	want := false

	assert.Equal(t, want, Expression{root: &node}.Evaluate(ValueSet{"testTrue": struct{}{}}))
}
