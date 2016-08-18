package goqueue

import "fmt"

//IllegalTypeError can be returned when you make an illegal comparisson.
//Mainly exported to allow the possibility of specific error logic.
type IllegalTypeError struct {
	expected string
	actual   string
}

func (e IllegalTypeError) Error() string {
	return fmt.Sprintf("Illegal comparisson of type %s to expected %s.", e.actual, e.expected)
}
