package bsfilter

// SymbolType is an enumerated value of all parseable symbols.
type SymbolType int

// A Symbol is a single kind of token.  It may be a named token or an
// operator.
type Symbol struct {
	t SymbolType
	ident string
}
