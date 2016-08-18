# goqueue

A simple priority queue implementation in golang.

As of now this is a minqueue (i.e. the smallest element gets dequeue'd first,
as this is the reason why I wrote this / my use case. I'll probably generalise
this sometime down the road.)

There are now external dependencies apart from the stdlib.

I have not done any extensive performance tests. The underlying structure is based on a binary heap,
which should give very decent results.


There is no reason not to use this software in the sense that it works and is stable, but 
I do not know if I will keep maintaining this on a regular schedule. Please be aware of that.

Contributions with regards to the above mentioned limitations (or lack of evidence)
is welcome.


## License

This software is licensed under the permissive MIT license.

Consult the [license file](LICENSE.MD) for more information.

## Usage & Example

* Import the library.

* Define an ```Orderable``` type

```go
type ordInt struct {
    x int
} 

func (c ordInt) compareTo(other Orderable) (int, error) {

	o, err := other.(ordInt)

	switch {
	case !err:
		return 2, &IllegalTypeError{actual: reflect.TypeOf(other).Name(), expected: "ordInt"}
	case c.x > o.x:
		return 1, nil
	case c.x < o.x:
		return -1, nil
	default:
		return 0, nil
	}
}
```

* Enqueue and dequeue, be on your merry way

```go
q := goqueue.New()
q.enqueue(ordInt{24}, anyObjectMightBeStoredHere)
q.enqueue(ordInt{23}, anyObjectMightBeStoredHere)

q.dequeue() //Will result in 23's object.
```

## Tests

Tests can simply be run by ```go test```

## Version history

Semver is used.

### 0.1.0
Basic prio(min)queue functionality implemented and functionaly tested.
