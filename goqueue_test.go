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

	switch {
	case !err:
		return 2, &IllegalTypeError{actual: reflect.TypeOf(other).Name(), expected: "compInt"}
	case c.x > o.x:
		return 1, nil
	case c.x < o.x:
		return -1, nil
	default:
		return 0, nil
	}
}

//Super quick compInt constructor.
func ci(x int) compInt {
	return compInt{x}
}

type compString struct {
	x string
}

func (c compString) compareTo(other Orderable) (int, error) {

	o, err := other.(compString)

	switch {
	case !err:
		return 2, &IllegalTypeError{actual: reflect.TypeOf(other).Name(), expected: "compString"}
	case c.x > o.x:
		return 1, nil
	case c.x < o.x:
		return -1, nil
	default:
		return 0, nil
	}
}

func TestEnqueueShouldWorkOnEmptyQueue(t *testing.T) {

	queue := New()

	prio := ci(70)
	val := true
	queue.Enqueue(prio, val)

	if queue.items[0].v != val || queue.items[0].k != prio {
		t.Fatalf("Expected to enqueue %t at prio %d, found %+v", true, prio, queue.items[0])
	}

}

func TestEnqueueShouldGiveAnErrorIfYouTryToEnterAnIllegalType(t *testing.T) {

	queue := New()

	p1 := ci(1)
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

//validateHeapProperty validates the entire queue's heap property.
//(In probably a suboptimal manner, but it'll do for this tests.)
func validateHeapProperty(queue goQueue) error {

	qLen := len(queue.items)

	if len(queue.items) <= 1 {
		return nil
	}

	for i, item := range queue.items {
		//The root item has no parent. We'll apply a different check.
		if i == 0 {
			continue
		}

		itemParent := queue.items[computeParentIdx(i)]
		if comp, err := item.k.compareTo(itemParent.k); err == nil && comp < 0 {
			return fmt.Errorf("Heap property violated: item %+v had parent %+v.", item, itemParent)
		}
	}

	//Checks for the root element
	lIdx, rIdx := computeChildIndices(0)
	if lIdx <= qLen {
		if comp, err := queue.items[0].k.compareTo(queue.items[lIdx].k); err == nil && comp > 0 {
			return fmt.Errorf("Heap property violated: root %+v had child %+v", queue.items[0], queue.items[lIdx])
		}

	}
	if rIdx <= qLen {
		if comp, err := queue.items[0].k.compareTo(queue.items[rIdx].k); err == nil && comp > 0 {
			return fmt.Errorf("Heap property violated: root %+v had child %+v", queue.items[0], queue.items[rIdx])
		}
	}

	return nil
}

func TestBalancingEnqueue(t *testing.T) {

	queue := New()
	//Enqueue some values with different prios.
	queue.Enqueue(ci(5), 0)
	queue.Enqueue(ci(4), 0)
	queue.Enqueue(ci(3), 0)
	queue.Enqueue(ci(2), 0)
	queue.Enqueue(ci(100), 0)
	queue.Enqueue(ci(6), 0)
	queue.Enqueue(ci(6), 0)

	/*
		This is the expected bin heap
		(variations possible, but primary property is that no key must be higher than it's parent.)

				1
			2/		\3
		4/	   \5|6/

	*/

	if valid := validateHeapProperty(*queue); valid != nil {
		t.Fatal(valid)
	}

}

func TestEmptyQueueDequeue(t *testing.T) {
	queue := New()

	item, _ := queue.Dequeue()

	if item != nil {
		t.Fatalf("We expected to get nil from dequeue'ing an empty queue.")
	}
}

func TestFilledQueueDequeue(t *testing.T) {
	queue := New()

	queue.Enqueue(ci(50), 5)
	queue.Enqueue(ci(20), 2)
	queue.Enqueue(ci(40), 4)
	queue.Enqueue(ci(600), 6)
	queue.Enqueue(ci(30), 3)
	queue.Enqueue(ci(10), 1)

	expectedOutput := [6]int{1, 2, 3, 4, 5, 6}

	for i, expected := range expectedOutput {
		val, _ := queue.Dequeue()
		if val != expected {
			t.Fatalf("Expected to retrieve %d at index %d got %+v", expected, i, val)
		}
	}
}
