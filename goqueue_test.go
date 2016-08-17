package goqueue

import (
	"fmt"
	"reflect"
	"testing"
)

type compInt struct {
	x int
}

func (c compInt) compareTo(other Orderable) (int, error) {

	o, err := other.(compInt)

	if !err {
		return 2, &IllegalTypeError{actual: reflect.TypeOf(other).Name(), expected: "compInt"}
	}

	if c.x > o.x {
		return 1, nil
	} else if c.x < o.x {
		return -1, nil
	}

	return 0, nil
}

type compString struct {
	x string
}

func (c compString) compareTo(other Orderable) (int, error) {

	o, err := other.(compString)

	if !err {
		return 2, &IllegalTypeError{actual: reflect.TypeOf(other).Name(), expected: "compString"}
	}

	if c.x > o.x {
		return 1, nil
	} else if c.x < o.x {
		return -1, nil
	}

	return 0, nil

}

func TestEnqueueShouldWorkOnEmptyQueue(t *testing.T) {

	queue := New()

	prio := compInt{70}
	val := true
	queue.Enqueue(prio, val)

	if queue.items[0].v != val || queue.items[0].k != prio {
		t.Fatalf("Expected to enqueue %t at prio %d, found %+v", true, prio, queue.items[0])
	}

}

func TestEnqueueShouldGiveAnErrorIfYouTryToEnterAnIllegalType(t *testing.T) {

	queue := New()

	p1 := compInt{1}
	v1 := true

	p2 := compString{"somestr"}
	v2 := false

	err := queue.Enqueue(p1, v1)

	if err != nil {
		t.Fail()
	}

	err = queue.Enqueue(p2, v2)

	if _, ok := err.(*IllegalTypeError); !ok {
		t.Fatalf("Expected an IllegalTypeError because we initialized with" +
			"a compInt and tried to insert a compString.")
	}

}

func TestBalancingEnqueue(t *testing.T) {

	queue := New()

	p1 := compInt{5}
	p2 := compInt{4}
	p3 := compInt{3}
	p4 := compInt{2}
	p5 := compInt{1}
	p6 := compInt{10}

	queue.Enqueue(p5, 0)
	queue.Enqueue(p1, 0)
	queue.Enqueue(p2, 0)
	queue.Enqueue(p3, 0)
	queue.Enqueue(p4, 0)
	queue.Enqueue(p6, 0)

	fmt.Println()
	fmt.Println()
	for _, i := range queue.items {
		fmt.Println(i)
	}

	t.Fatalf("You should implement a test")
}
