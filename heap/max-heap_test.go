package heap

import (
	"math/rand"
	"slices"
	"testing"
)

type Item struct {
	name     string
	priority int
}

// using int keys. MaxHeap should place highest K on top

// maxcmp returns a higher value if a is greater than b
func maxcmp(a int, b int) int {
	return a - b
}

func newint(i int) *int {
	x := i
	return &x
}

func Test_heapifySimple(t *testing.T) {
	h := NewHeap[int](maxcmp)
	h.a[0] = 1
	h.a[1] = 2
	h.a[2] = 3
	h.heapsize = 3

	h.heapify(0)

	if h.a[0] != 3 {
		t.Errorf("error in heapify, got %d", h.a[0])
	}
	if !validateHeapProperty(h) {
		t.Errorf("max heap property not satisfied")
	}
}

func Test_RandomBuildMaxHeap(t *testing.T) {
	r := rand.New(rand.NewSource(123))
	var a []int
	for i := 0; i < 1000000; i++ {
		a = append(a, r.Int())
	}

	h := BuildHeap(a, maxcmp)

	if !validateHeapProperty(h) {
		t.Errorf("error, max heap property is not maintained!")
	}
}

func Test_RandomHeapsort(t *testing.T) {
	r := rand.New(rand.NewSource(123))
	var a []int
	for i := 0; i < 1000000; i++ {
		a = append(a, r.Int())
	}

	Heapsort(a, maxcmp)

	if !slices.IsSortedFunc(a, maxcmp) {
		t.Errorf("error, heapsort didnt!")
	}
}

func TestHeap_Extract(t *testing.T) {

	var a []int
	r := rand.New(rand.NewSource(123))
	//insert 1 million random ints
	for i := 0; i < 1000000; i++ {
		a = append(a, r.Int())
	}
	h := BuildHeap(a, maxcmp)
	//extract each one and test that it is smaller than the previous
	prev, _ := h.Extract()
	for i := 1; i < 1000000; i++ { // start at one, I already removed one.
		curr, _ := h.Extract()
		if prev < curr {
			t.Errorf("Error, inserted objects out of order. prev = %d, curr = %d", prev, curr)
		}
		prev = curr
	}
}

func validateHeapProperty[V any](h *Heap[V]) bool {
	for i := h.heapsize - 1; i >= 0; i-- {
		if h.cmp(h.a[i], h.a[h.parent(i)]) > 0 {
			return false
		}
	}
	return true
}
