package rbtree

import (
	"math/rand"
	"slices"
	"strings"
	"testing"
)

func BenchmarkRBTree_Insert(b *testing.B) {

	tree := New[int, int](func(a int, b int) int {
		return a - b
	})
	r := rand.New(rand.NewSource(123))
	duplicates := 0

	for i := 0; i < b.N; i++ {
		obj := r.Int()
		_, dup := tree.Insert(obj, obj)
		if dup {
			duplicates++
		}
	}

	//fmt.Printf("<<<%d>>>", tree.Height())
}

// this is mega fast?
func BenchmarkRBTree_InOrderInsert(b *testing.B) {

	tree := New[int, int](func(a int, b int) int {
		return a - b
	})
	duplicates := 0

	for i := 0; i < b.N; i++ {
		_, dup := tree.Insert(i, i)
		if dup {
			duplicates++
		}
	}

	//fmt.Printf("<<<%d>>>", tree.Height())
}

func intcmp(a int, b int) int {
	return a - b
}

func newString(x string) *string {
	y := x
	return &y
}

func newInt(x int) *int {
	y := x
	return &y
}
func TestBST_Insert(t *testing.T) {
	b := New[int, string](intcmp)
	b.Insert(5, "5")
	b.Insert(3, "3")
	b.Insert(4, "4")
	b.Insert(1, "1")

	x := valueSlice(b)

	if !slices.IsSortedFunc(x, strings.Compare) {
		t.Errorf("inserts were not sorted!")
	}

	if !slices.Contains(x, "1") {
		t.Errorf("missing 1")
	}

	if !slices.Contains(x, "3") {
		t.Errorf("missing 3")
	}

	if !slices.Contains(x, "4") {
		t.Errorf("missing 4")
	}

	if !slices.Contains(x, "5") {
		t.Errorf("missing 5")
	}
}

func TestBST_InsertDup(t *testing.T) {
	b := New[int, string](intcmp)

	b.Insert(5, "5")
	b.Insert(3, "3")
	b.Insert(4, "4")
	b.Insert(1, "1")

	got, _ := b.Insert(4, "166")
	if got != "4" {
		t.Errorf("inserting a duplicate, should see previous value 4, got %s", got)
	}

	got, _ = b.Search(4)

	if got != "166" {
		t.Errorf("duplicate Key K was replaces with 166, got %s", got)
	}
}

func TestBST_Predecessor(t *testing.T) {
	// pred can be max on min side or above in the tree

	//build a slice starting at maximum and call pred
	//call when there is no children

	b := New[int, int](intcmp)

	b.Insert(15, 15)
	b.Insert(6, 6)
	b.Insert(3, 3)
	b.Insert(2, 2)
	b.Insert(4, 4)
	b.Insert(7, 7)
	b.Insert(13, 13)
	b.Insert(9, 9)
	b.Insert(18, 18)
	b.Insert(17, 17)
	b.Insert(20, 20)

	_, found := b.Predecessor(2)
	if found {
		t.Errorf("there should be no predecessor of min")
	}
	got, _ := b.Predecessor(15)
	if got != 13 {
		t.Errorf("didn't find max of min side, got %d", got)
	}
	got, _ = b.Predecessor(17)
	if got != 15 {
		t.Errorf("didn't traverse up properly, got %d", got)
	}
}

func TestBST_Successor(t *testing.T) {
	// succ can be min on max side or above in the tree

	//build a slice starting at minimum and call succ
	//call when there is no children

	b := New[int, int](intcmp)

	b.Insert(15, 15)
	b.Insert(6, 6)
	b.Insert(3, 3)
	b.Insert(2, 2)
	b.Insert(4, 4)
	b.Insert(7, 7)
	b.Insert(13, 13)
	b.Insert(9, 9)
	b.Insert(18, 18)
	b.Insert(17, 17)
	b.Insert(20, 20)

	got, found := b.Successor(20)
	if found {
		t.Errorf("there should be no successor of max, got %d", got)
	}
	got, found = b.Successor(15)
	if got != 17 {
		t.Errorf("didn't find min of max side, got %d", got)
	}
	got, found = b.Successor(13)
	if got != 15 {
		t.Errorf("didn't traverse up properly, got %d", got)
	}
}

func TestBST_Clear(t *testing.T) {
	b := New[int, int](intcmp)

	b.Insert(15, 15)
	b.Insert(6, 6)
	b.Insert(3, 3)
	b.Insert(2, 2)
	b.Insert(4, 4)
	b.Insert(7, 7)
	b.Insert(13, 13)
	b.Insert(9, 9)
	b.Insert(18, 18)
	b.Insert(17, 17)
	b.Insert(20, 20)

	b.Clear()

	if b.Size() != 0 {
		t.Errorf("after clear there should be no nodes! got %d", b.Size())
	}

	if b.root != nil {
		t.Errorf("t.root is not nil")
	}
}

func TestBST_Size(t *testing.T) {

	b := New[int, string](intcmp)

	b.Insert(5, "5")
	b.Insert(3, "3")
	b.Insert(4, "4")
	b.Insert(1, "1")

	if b.Size() != 4 {
		t.Errorf("incorrect size, expected 4, got %d", b.Size())
	}

	b.Remove(3)

	if b.Size() != 3 {
		t.Errorf("incorrect size, expected 3, got %d", b.Size())
	}
}

func TestBST_Height(t *testing.T) {
	b := New[int, int](intcmp)

	b.Insert(15, 15)
	b.Insert(6, 6)
	b.Insert(3, 3)
	b.Insert(2, 2)
	b.Insert(4, 4)
	b.Insert(7, 7)
	b.Insert(13, 13)
	b.Insert(9, 9)
	b.Insert(18, 18)
	b.Insert(17, 17)
	b.Insert(20, 20)

	if b.Height() != 5 {
		t.Errorf("Height was expected 5, got %d", b.Height())
	}

	b.Remove(9)

	if b.Height() != 4 {
		t.Errorf("Height was expected 4, got %d", b.Height())
	}
}

func TestBST_Remove(t *testing.T) {
	// four cases to consider. No right child, no left child, and both children (2 cases)

	b := New[int, int](intcmp)

	b.Insert(15, 15)
	b.Insert(6, 6)
	b.Insert(3, 3)
	b.Insert(2, 2)
	b.Insert(4, 4)
	b.Insert(7, 7)
	b.Insert(13, 13)
	b.Insert(9, 9)
	b.Insert(18, 18)
	b.Insert(17, 17)
	b.Insert(20, 20)

	//case 1: delete 13 (replaced with 9)
	got, _ := b.Remove(13)
	if got != 13 {
		t.Errorf("13 not returned by remove")
	}

	x := valueSlice(b)

	if !slices.IsSortedFunc(x, intcmp) {
		t.Errorf("tree structure incorrect")
	}
	if slices.Contains(x, 13) {
		t.Errorf("13 not removed")
	}

	//case 2: delete 2 (no children) then delete 3

	got, _ = b.Remove(2)

	if got != 2 {
		t.Errorf("2 not returned by remove")
	}
	got, _ = b.Remove(3)

	if got != 3 {
		t.Errorf("3 not returned by remove")
	}

	x = valueSlice(b)

	if !slices.IsSortedFunc(x, intcmp) {
		t.Errorf("tree structure incorrect")
	}
	if slices.Contains(x, 3) {
		t.Errorf("3 not removed")
	}

	// case 3a: delete 17 then delete 15

	got, _ = b.Remove(17)
	if got != 17 {
		t.Errorf("17 not returned by remove")
	}

	got, _ = b.Remove(15)
	if got != 15 {
		t.Errorf("15 not returned by remove")
	}

	x = valueSlice(b)

	if !slices.IsSortedFunc(x, intcmp) {
		t.Errorf("tree structure incorrect")
	}
	if slices.Contains(x, 15) {
		t.Errorf("15 not removed")
	}

	// case 3b: insert 19, 25, 21, 23, 22 24, then delete 20

	b.Insert(19, 19)
	b.Insert(25, 25)
	b.Insert(21, 21)
	b.Insert(23, 23)
	b.Insert(22, 22)
	b.Insert(24, 24)

	got, _ = b.Remove(20)
	if got != 20 {
		t.Errorf("20 not returned by remove")
	}

	x = valueSlice(b)

	if !slices.IsSortedFunc(x, intcmp) {
		t.Errorf("tree structure incorrect")
	}
	if slices.Contains(x, 20) {
		t.Errorf("20 not removed")
	}

	if b.Size() != 11 {
		t.Errorf("Size incorrect! expected 11 but was %d", b.Size())
	}

	// finally, remove something not in the tree

	_, found := b.Remove(2500)
	if found {
		t.Errorf("error, should return nil for keys not in the tree")
	}
}

func valueSlice[K any, V any](b *RBTree[K, V]) []V {
	var x []V
	b.traverse(b.root, func(n *node[K, V]) {
		x = append(x, n.val)
	})
	return x
}

func keySlice[K any, V any](b *RBTree[K, V]) []K {
	var x []K
	b.traverse(b.root, func(n *node[K, V]) {
		x = append(x, n.key)
	})
	return x
}
