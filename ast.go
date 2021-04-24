package bsfilter

// Evaluate runs the expression against the supplied ValueSet
func (e Expression) Evaluate(vset ValueSet) bool {
	return e.root.Resolve(vset)
}

func (e Expression) String() string {
	return e.root.String()
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
