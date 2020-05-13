package myint

import "errors"

var ErrorDivideByZero = errors.New("should never divide by zero")

type MyInt int32

// Add is adding to Myint.
func (i *MyInt) Add(n int) {
	*i += MyInt(n)
}

func (i MyInt) Sub(n int) MyInt {
	return i - MyInt(n)
}

func (i MyInt) Multiply(n int) MyInt {
	return i * MyInt(n)
}

func (i MyInt) Divide(n int) (MyInt, error) {
	if n == 0 {
		return i, ErrorDivideByZero
	}
	return i / MyInt(n), nil
}
