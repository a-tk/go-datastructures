package heap

import (
	"math/rand"
	"testing"
)

func BenchmarkHeap_Insert(b *testing.B) {
	h := NewPriorityQueue[int](maxcmp)
	r := rand.New(rand.NewSource(123))
	for i := 0; i < b.N; i++ {
		obj := r.Int()
		h.Insert(obj)
	}
}
func BenchmarkHeap_RandomInsertExtract(b *testing.B) {

	h := NewPriorityQueue[int](maxcmp)
	r := rand.New(rand.NewSource(123))

	for i := 0; i < b.N; i++ {
		obj := r.Int()
		h.Insert(obj)
	}
	//extract each one
	for i := 0; i < b.N; i++ {
		h.Extract()
	}
}

func Test_InsertMaxSimple(t *testing.T) {

	pq := NewPriorityQueue[int](maxcmp)
	pq.Insert(1)
	pq.Insert(2)
	pq.Insert(3)

	got, _ := pq.Extract()

	if got != 3 {
		t.Errorf("expected 3, got %d", got)
	}

	got, _ = pq.Extract()

	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
	got, _ = pq.Extract()
	if got != 1 {
		t.Errorf("expected 1, got %d", got)
	}
}

func Test_UpdateLarger(t *testing.T) {

	h := NewPriorityQueue[int](maxcmp)
	one := 1
	two := 2
	three := 3
	four := 4
	h.Insert(one)
	h.Insert(two)
	h.Insert(three)
	h.Insert(four)

	got, _ := h.Top()

	if got != 4 {
		t.Errorf("expected 4, got %d", got)
	}

	//TODO: support updating lower or higher
	//six := 6
	//
	//h.Update(three, six)
	//got = h.Top()
	//
	//if *got != 6 {
	//	t.Errorf("expected 6, got %d", *got)
	//}
	//
	//h.Update(six, three)
	//got = h.Top()
	//
	//if *got != 4 {
	//	t.Errorf("expected 4, got %d", *got)
	//}
}

func TestHeap_RandomInsertExtract(t *testing.T) {

	h := NewPriorityQueue[int](maxcmp)
	r := rand.New(rand.NewSource(123))
	//insert 1 million random ints
	for i := 0; i < 1000000; i++ {
		obj := r.Int()
		h.Insert(obj)
	}
	if !validatePQProperty(h) {
		t.Errorf("MaxHeapProperty not observed")
	}
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

func TestHeap_Item(t *testing.T) {
	itemA := Item{
		"itemA",
		1,
	}
	itemB := Item{
		"itemB",
		200,
	}
	itemC := Item{
		"itemC",
		5,
	}

	h := NewPriorityQueue[Item](itemCmp)

	h.Insert(itemA)
	h.Insert(itemB)
	h.Insert(itemC)

	if !validatePQProperty(h) {
		t.Errorf("MaxHeapProperty not observed")
	}
}

func itemCmp(a Item, b Item) int {
	return a.priority - b.priority
}

func validatePQProperty[V any](h *PriorityQueue[V]) bool {
	for i := h.heapsize - 1; i >= 0; i-- {
		if h.cmp(h.a[i], h.a[h.parent(i)]) > 0 {
			return false
		}
	}
	return true
}
