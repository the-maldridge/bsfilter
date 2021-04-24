package bsfilter

import "errors"

var (
	// ErrInvalidExpression is returned when an expression cannot
	// be constructed based on the input data.
	ErrInvalidExpression = errors.New("supplied expression is invalid")
)
