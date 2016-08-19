package goqueue

import "math"

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

//computeParentIdx computes the index of a parent given a childIndex.
func computeParentIdx(childIdx int) int {
	return int(math.Floor(float64(childIdx-1) / 2))
}

//computeChildIndices computes the indices of the childs.
//Those are returned left, right (as you'd expect.)
func computeChildIndices(childIdx int) (int, int) {
	return 2*childIdx + 1, 2*childIdx + 2
}

//Enqueue adds a new item to the queue, with a given priority.
func (q *goQueue) Enqueue(prio Orderable, item interface{}) error {

	if len(q.items) == 0 {
		q.items = append(q.items, goQueueItem{k: prio, v: item})
		return nil
	}

	//Type check the prio.
	someRandomItem := q.items[0]
	_, comperr := someRandomItem.k.CompareTo(prio)
	if comperr != nil {
		return comperr
	}

	q.items = append(q.items, goQueueItem{k: prio, v: item})
	q.percUp()

	return nil
}

func (q *goQueue) percUp() {
	qLen := len(q.items)

	//Nothing to balance? Nice. Be gone.
	if qLen <= 1 {
		return
	}

	idx := qLen - 1
	item := q.items[idx]
	pIdx := computeParentIdx(idx)
	p := q.items[pIdx]

	comp, _ := item.k.CompareTo(p.k)

	for idx != 0 && comp < 0 {
		q.items[pIdx] = item
		q.items[idx] = p

		idx = pIdx
		if idx == 0 {
			break
		}
		pIdx = computeParentIdx(idx)
		p = q.items[pIdx]

		comp, _ = item.k.CompareTo(p.k)
	}
}

func (q *goQueue) Dequeue() (interface{}, Orderable) {
	if len(q.items) == 0 {
		return nil, nil
	}

	elem := q.items[0]

	//1. Insert last item in place of the dequeue'd.
	item := q.items[len(q.items)-1]
	q.items[0] = item
	//2. Remove now double present last item.
	q.items = q.items[1:]

	q.percDown(0)
	return elem.v, elem.k
}

func (q *goQueue) percDown(itemIdx int) {

	qLen := len(q.items)

	//Nothing to do anymore. The queue is empty.
	if qLen <= 1 {
		return
	}

	item := q.items[itemIdx]

	//Now check if the heap property is still valid.
	lIdx, rIdx := computeChildIndices(0)

	smallest := itemIdx
	if lIdx < qLen {
		if comp, err := item.k.CompareTo(q.items[lIdx].k); err == nil && comp >= 0 {
			smallest = lIdx
		}
	}
	if rIdx < qLen {
		if comp, err := item.k.CompareTo(q.items[rIdx].k); err == nil && comp >= 0 {
			smallest = rIdx
		}
	}
	if smallest != itemIdx {
		largestItem := q.items[smallest]
		q.items[smallest] = item
		q.items[itemIdx] = largestItem
		q.percDown(smallest)
	}
}
