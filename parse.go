package bsfilter

import (
	"log"
	"strings"
	"unicode"
)

func (t SymbolType) String() string {
	return [...]string{"LPAREN", "RPAREN", "AND", "OR", "NOT", "IDENT"}[t]
}

// Tokenize splits the input string into distinct tokens for further
// processing.
func Tokenize(in string) []Symbol {
	idents := strings.FieldsFunc(in, func(r rune) bool {
		return strings.ContainsAny(string(r), "()&|!")
	})

	p := ""
	for _, r := range in {
		if unicode.IsSpace(r) {
			continue
		}
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if len(p) > 0 && string(p[len(p)-1]) == "I" {
				continue
			}
			p += "I"
		} else {
			p += string(r)
		}
	}

	out := []Symbol{}
	identCtr := 0
	for _, r := range p {
		switch r {
		case '(':
			out = append(out, Symbol{T: SymbolLParen})
		case ')':
			out = append(out, Symbol{T: SymbolRParen})
		case '&':
			out = append(out, Symbol{T: SymbolBinaryAnd})
		case '|':
			out = append(out, Symbol{T: SymbolBinaryOr})
		case '!':
			out = append(out, Symbol{T: SymbolUnaryNot})
		case 'I':
			out = append(out, Symbol{T: SymbolIdent, Ident: strings.TrimSpace(idents[identCtr])})
			identCtr++
		}
	}
	return out
}

// New returns an expression that is ready to be parsed or an error if
// the expression cannot be successfully parsed.
func New(in string) (*Expression, error) {
	tokens := Tokenize(in)
	if len(tokens) == 0 {
		return nil, ErrInvalidExpression
	}
	return NewFromTokens(tokens), nil
}

// NewFromTokens provides a way to construct the entire
func NewFromTokens(tokens []Symbol) *Expression {
	p := &Parser{
		symbols: tokens,
		l:       log.Default(),
	}

	p.expression()

	e := &Expression{l: log.Default(), root: p.root}

	return e
}

func (p *Parser) expression() {
	p.term()
	p.l.Printf("[EXPRESSION] %+v", p.curSymbol)
	for p.curSymbol.T == SymbolBinaryOr {
		or := &ASTNode{Symbol: p.curSymbol}
		or.Left = p.root
		p.term()
		or.Right = p.root
		p.root = or
	}
}

func (p *Parser) term() {
	p.factor()
	p.l.Printf("[TERM] %+v", p.curSymbol)
	for p.curSymbol.T == SymbolBinaryAnd {
		and := &ASTNode{Symbol: p.curSymbol}
		and.Left = p.root
		p.term()
		and.Right = p.root
		p.root = and
	}
}

func (p *Parser) factor() {
	p.nextSymbol()
	p.l.Printf("[FACTOR] %+v", p.curSymbol)
	switch p.curSymbol.T {
	case SymbolIdent:
		p.root = &ASTNode{Symbol: p.curSymbol}
		p.nextSymbol()
	case SymbolUnaryNot:
		n := &ASTNode{Symbol: p.curSymbol}
		p.factor()
		n.Right = p.root
		p.root = n
	case SymbolLParen:
		p.expression()
		p.nextSymbol() // Discard the SymbolRParen
	}
}

func (p *Parser) nextSymbol() {
	if p.curPos >= len(p.symbols) {
		return
	}
	p.curSymbol = &p.symbols[p.curPos]
	p.l.Printf("__SYMBOL__: %s", p.curSymbol)
	p.curPos++
}
