package goqueue

//Orderable is an interface that defines types that are compareable and can yield an inequality.
//I.e. an Integer naturally has this property because we can say: int a > int b. Note that you
//have to create a custom type to make this happen (boxing int in a struct.) There is now monkey
//patching in Go.
//
//With regards to the error: you can throw an IllegalTypeError when a comparisson happens
//between two distinct types (implementors of this interface) that have no solution to their
//equality question. (For performance reasons, do this as early as possible.)
type Orderable interface {
	CompareTo(other Orderable) (int, error)
}
