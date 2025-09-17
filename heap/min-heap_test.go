package heap

import "testing"

// using int keys. Min heap should put min val on top

// mincmp returns a higher value if a is greater than b
func mincmp(a int, b int) int {
	return b - a
}

func Test_heapifyMinSimple(t *testing.T) {
	h := NewHeap[int](mincmp)
	h.a[0] = 3
	h.a[1] = 2
	h.a[2] = 1
	h.heapsize = 3

	h.heapify(0)

	if h.a[0] != 1 {
		t.Errorf("error in heapify, got %d", h.a[0])
	}
}
