package goqueue

import "fmt"

type IllegalTypeError struct {
	expected string
	actual   string
}

func (e IllegalTypeError) Error() string {
	return fmt.Sprintf("Illegal comparisson of type %s to expected %s.", e.actual, e.expected)
}
