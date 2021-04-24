package bsfilter

import "log"

// SymbolType is an enumerated value of all parseable symbols.
type SymbolType int

// A Symbol is a single kind of token.  It may be a named token or an
// operator.
type Symbol struct {
	T     SymbolType
	Ident string
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

// An Expression is a single boolean expression that is parsed and has
// an evaluator attached.
type Expression struct {
	l *log.Logger

	root *ASTNode
}

// A Parser creates a new expression out of a string by first
// tokenizing it and then recursing through the expression.
type Parser struct {
	l *log.Logger

	curSymbol *Symbol
	curPos    int
	symbols   []Symbol

	root *ASTNode
}

// An ASTNode is, unsurprisingly, a node on the abstract syntax tree
// that is parsed from the expression.
type ASTNode struct {
	*Symbol

	Left  *ASTNode
	Right *ASTNode
}

// A ValueSet is a set of attributes.  Any attribute that is not
// present in the set is implicitly false.
type ValueSet map[string]struct{}
