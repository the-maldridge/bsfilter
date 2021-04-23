package bsfilter

import (
	"fmt"
	"strings"
	"unicode"
)

func (t SymbolType) String() string {
	return [...]string{"(", ")", "&", "|", "!", "IDENT"}[t]
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
	fmt.Println(idents)
	fmt.Println(p)

	out := []Symbol{}
	identCtr := 0
	for _, r := range p {
		switch r {
		case '(':
			out = append(out, Symbol{t: SymbolLParen})
		case ')':
			out = append(out, Symbol{t: SymbolRParen})
		case '&':
			out = append(out, Symbol{t: SymbolBinaryAnd})
		case '|':
			out = append(out, Symbol{t: SymbolBinaryOr})
		case '!':
			out = append(out, Symbol{t: SymbolUnaryNot})
		case 'I':
			out = append(out, Symbol{t: SymbolIdent, ident: strings.TrimSpace(idents[identCtr])})
			identCtr++
		}
	}
	return out
}
