package heap

type Heap[E any] struct {
	a        []E
	heapsize int
	cmp      func(E, E) int
}

func NewHeap[E any](cmp func(E, E) int) *Heap[E] {
	return &Heap[E]{
		heapsize: 0,
		cmp:      cmp,
		a:        make([]E, 8),
	}
}

func (h *Heap[E]) left(i int) int {
	return 2*i + 1
}

func (h *Heap[E]) right(i int) int {
	return 2*i + 2
}

func (h *Heap[E]) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap[E]) heapify(i int) {
	l := h.left(i)
	r := h.right(i)
	var z int
	if l < h.heapsize && h.cmp(h.a[i], h.a[l]) < 0 {
		z = l
	} else {
		z = i
	}
	if r < h.heapsize && h.cmp(h.a[z], h.a[r]) < 0 {
		z = r
	}
	if z != i {
		h.swap(i, z)
		h.heapify(z)
	}
}

func (h *Heap[E]) Top() (top E, ok bool) {
	if h.heapsize < 1 {
		return top, false
	} else {
		return h.a[0], true
	}
}

// consider nil-ing elements for the garbage collector to consume
func (h *Heap[E]) Extract() (val E, ok bool) {
	v, ok := h.Top()
	if !ok {
		return val, false
	}
	h.a[0] = h.a[h.heapsize-1]
	h.a = h.a[:h.heapsize-1] // free for garbage collection //won't work with nil unless I force pointers in V
	h.heapsize--
	h.heapify(0)
	return v, true
}

func (h *Heap[E]) swap(i int, j int) {
	t := h.a[i]
	h.a[i] = h.a[j]
	h.a[j] = t
}

func BuildHeap[E any](a []E, cmp func(E, E) int) *Heap[E] {
	h := NewHeap(cmp)
	h.heapsize = len(a)
	h.a = a
	for i := (h.heapsize - 1) / 2; i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func Heapsort[E any](a []E, cmp func(E, E) int) {
	h := BuildHeap(a, cmp)
	for i := len(a) - 1; i >= 1; i-- {
		h.swap(0, i)
		h.heapsize--
		h.heapify(0)
	}
}
