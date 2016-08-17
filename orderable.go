package goqueue

type Orderable interface {
	compareTo(other Orderable) (int, error)
}
