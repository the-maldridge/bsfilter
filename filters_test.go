package bsfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressionSet(t *testing.T) {
	cases := []struct {
		vset ValueSet
		want []string
	}{
		{
			vset: ValueSet{
				"foo": struct{}{},
			},
			want: []string{"expr2"},
		},
		{
			vset: ValueSet{
				"foo": struct{}{},
				"bar": struct{}{},
			},
			want: []string{"expr1"},
		},
		{
			vset: ValueSet{
				"foo": struct{}{},
				"baz": struct{}{},
			},
			want: []string{"expr2", "expr3", "expr4"},
		},
		{
			vset: ValueSet{
				"foo": struct{}{},
				"bar": struct{}{},
				"baz": struct{}{},
			},
			want: []string{"expr1", "expr3"},
		},
	}

	exprs := map[string]string{
		"expr1": "foo&bar",
		"expr2": "foo&!bar",
		"expr3": "baz",
		"expr4": "foo&!bar&baz",
	}

	es := NewExpressionSet()

	for n, e := range exprs {
		ex, _ := New(e)
		es.Add(n, ex)
	}

	for _, c := range cases {
		assert.ElementsMatch(t, c.want, es.Filter(c.vset))
	}

	vset := ValueSet{"foo": struct{}{}, "bar": struct{}{}, "baz": struct{}{}}
	es.Del("expr3")
	assert.ElementsMatch(t, []string{"expr1"}, es.Filter(vset))
}
