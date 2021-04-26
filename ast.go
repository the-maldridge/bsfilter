package bsfilter

// Evaluate runs the expression against the supplied ValueSet
func (e Expression) Evaluate(vset ValueSet) bool {
	return e.root.Resolve(vset)
}

func (e Expression) String() string {
	return e.root.String()
}

// FilterValues takes a map of identifier to ValueSet and applies the
// expression to it outputing a list of any identifiers that had true
// evaluations.
func (e Expression) FilterValues(vset map[string]ValueSet) []string {
	out := []string{}
	for n, vs := range vset {
		if e.Evaluate(vs) {
			out = append(out, n)
		}
	}
	return out
}

// Resolve resolves an ASTNode to a boolean value.  By performing this
// recursion recursively we can resolve an entire AST to a single
// value.
func (a ASTNode) Resolve(vals ValueSet) bool {
	switch a.T {
	case SymbolUnaryNot:
		return !a.Right.Resolve(vals)
	case SymbolIdent:
		_, val := vals[a.Ident]
		return val
	case SymbolBinaryOr:
		return a.Left.Resolve(vals) || a.Right.Resolve(vals)
	case SymbolBinaryAnd:
		return a.Left.Resolve(vals) && a.Right.Resolve(vals)
	default:
		return false
	}
}

// String turns the expression into a string, mostly useful for
// debugging
func (a ASTNode) String() string {
	switch a.T {
	case SymbolUnaryNot:
		return "!" + a.Right.String()
	case SymbolIdent:
		return a.Ident
	case SymbolBinaryOr:
		return "(" + a.Left.String() + "|" + a.Right.String() + ")"
	case SymbolBinaryAnd:
		return "(" + a.Left.String() + "&" + a.Right.String() + ")"
	default:
		return ""
	}
}
