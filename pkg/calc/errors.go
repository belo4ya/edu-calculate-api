package calc

import "errors"

var (
	ErrInvalidExpr    = errors.New("invalid expression")
	ErrDivisionByZero = errors.New("division by zero")
)
