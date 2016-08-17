package goqueue

import (
	"fmt"
	"math"
)

type goQueue struct {
	items []goQueueItem
}

//New instantiates a new empty goQueue.
func New() *goQueue {
	q := new(goQueue)
	return q
}

type goQueueItem struct {
	k Orderable
	v interface{}
}

func computeParentIdx(childIdx int) int {
	return int(math.Floor(float64(childIdx-1) / 2))
}

func (q *goQueue) Enqueue(prio Orderable, item interface{}) error {

	if len(q.items) == 0 {
		q.items = append(q.items, goQueueItem{k: prio, v: item})
		return nil
	}

	//Type check the prio.
	someRandomItem := q.items[0]
	_, comperr := someRandomItem.k.compareTo(prio)
	if comperr != nil {
		return comperr
	}

	q.items = append(q.items, goQueueItem{k: prio, v: item})
	q.rebalance()

	return nil
}

func (q *goQueue) rebalance() {
	qLen := len(q.items)

	//Nothing to balance? Nice. Be gone.
	if qLen <= 1 {
		return
	}

	itemIdx := qLen - 1
	item := q.items[itemIdx]
	parentIdx := computeParentIdx(itemIdx)
	parent := q.items[parentIdx]

	comp, _ := item.k.compareTo(parent.k)

	for comp == -1 {
		fmt.Println(q.items)

		q.items[parentIdx] = item
		q.items[itemIdx] = parent
		itemIdx = parentIdx
		parentIdx = computeParentIdx(itemIdx)
		if parentIdx == -1 {
			break
		}
		parent = q.items[parentIdx]
		comp, _ = item.k.compareTo(parent.k)

	}

}

func (q *goQueue) DequeueMax() interface{} {
	//Simply return the last element.
	return nil
}

func (q *goQueue) DequeueMin() interface{} {
	//Simply return the first element.
	return nil
}
