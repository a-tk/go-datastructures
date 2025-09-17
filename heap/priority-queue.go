package heap

import "slices"

type PriorityQueue[E any] struct {
	*Heap[E]
}

func NewPriorityQueue[E any](cmp func(E, E) int) *PriorityQueue[E] {
	pq := &PriorityQueue[E]{
		Heap: NewHeap(cmp),
	}
	return pq
}

func (pq *PriorityQueue[E]) increaseKey(i int, newK E) (ok bool) {
	//if pq.cmp(newK, pq.a[i]) < 0 { // equal includes the use case from insert
	//	return false
	//}
	pq.a[i] = newK
	for i > 0 && pq.cmp(pq.a[i], pq.a[pq.parent(i)]) > 0 {

		pq.swap(i, pq.parent(i))
		i = pq.parent(i)
	}
	return true
}

func (pq *PriorityQueue[E]) Insert(e E) {

	//if len(pq.a) == pq.heapsize {
	//	pq.a = append(pq.a, e) // append will autogrow the slice if needed
	//} else {
	//	pq.a[pq.heapsize] = e
	//}

	pq.a = append(pq.a[:pq.heapsize], e)
	pq.heapsize++
	pq.increaseKey(pq.heapsize-1, e)
}

func (pq *PriorityQueue[E]) Update(old E, newKey E) (updated bool) {

	//find the index of element. O(n) :(
	index := slices.IndexFunc(pq.a, func(e E) bool {
		return pq.cmp(e, old) == 0
	})

	if index == -1 {
		return false
	}
	// update the key
	pq.a[index] = newKey

	if pq.cmp(old, newKey) < 0 {
		pq.heapify(index)
	} else if pq.cmp(old, newKey) > 0 {
		pq.increaseKey(index, newKey)
	}
	return true
}
