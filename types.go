package bsfilter

// SymbolType is an enumerated value of all parseable symbols.
type SymbolType int

// A Symbol is a single kind of token.  It may be a named token or an
// operator.
type Symbol struct {
	t SymbolType
	ident string
}

const (
	// SymbolLParen is a single left parenthesis '('
	SymbolLParen SymbolType = iota

	// SymbolRParen is a single right parenthesis ')'
	SymbolRParen

	// SymbolBinaryAnd is a binary and operator '&'
	SymbolBinaryAnd

	// SymbolBinaryOr is a binary or operator '|'
	SymbolBinaryOr

	// SymbolUnaryNot is a unary not operator '!'
	SymbolUnaryNot

	// SymbolIdent is a predicate identifier and may be an
	// arbitrary string of letters and numbers.
	SymbolIdent
)
