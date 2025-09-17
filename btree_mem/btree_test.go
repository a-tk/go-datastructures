package btree_mem

import (
	"math/rand"
	"slices"
	"testing"
)

func BenchmarkBTree_Insert(b *testing.B) {

	tree := New[int, int](20, func(a int, b int) int {
		return a - b
	})
	r := rand.New(rand.NewSource(123))
	duplicates := 0

	for i := 0; i < b.N; i++ {
		obj := r.Int()
		_, found := tree.Insert(obj, obj)
		if found {
			duplicates++
		}
	}

	//fmt.Printf("<<<%d>>>", tree.height)
	//fmt.Printf("<<<duplicates: %d>>>", duplicates)
}

func BenchmarkBTree_InOrderInsert(b *testing.B) {

	// degree 13 seems to be a good performance choice for in memory btree
	// benchmarks faster/op than the standard bst, but still 5x slower than built in map
	tree := New[int, int](13, func(a int, b int) int {
		return a - b
	})
	duplicates := 0

	for i := 0; i < b.N; i++ {
		_, found := tree.Insert(i, i)
		if found {
			duplicates++
		}
	}

	//fmt.Printf("<<<%d>>>", tree.height)
	//fmt.Printf("<<<duplicates: %d>>>", duplicates)
}

// TODO: read about go benchmarks
func Benchmark_BTreeInOrder(b *testing.B) {

	r := rand.New(rand.NewSource(123))

	tree := New[int, int](13, func(a int, b int) int {
		return a - b
	})

	for i := 0; i < b.N; i++ {
		for z := 0; z < 100000; z++ {
			obj := r.Int()
			tree.Insert(obj, obj)
		}

		var sorted []int = valueSlice(tree)

		slices.IsSorted(sorted)
	}
}

func Benchmark_BuiltinMap(b *testing.B) {

	r := rand.New(rand.NewSource(123))
	m := make(map[int]int)

	for i := 0; i < b.N; i++ {
		obj := r.Int()
		_, ok := m[obj]
		if !ok {
			m[obj] = obj
		}
	}
}

func Benchmark_BuiltinMapInOrder(b *testing.B) {

	r := rand.New(rand.NewSource(123))
	m := make(map[int]int)

	for i := 0; i < b.N; i++ {
		for z := 0; z < 100000; z++ {
			obj := r.Int()
			m[obj] = obj
		}

		var sorted []int

		for _, z := range m {
			sorted = append(sorted, z)
		}

		slices.Sort(sorted)
		slices.IsSorted(sorted)
	}
}

func newString(x string) *string {
	y := x
	return &y
}

func newInt(x int) *int {
	y := x
	return &y
}

func TestBTreeCreate(t *testing.T) {
	b := New[int, int](2, func(a int, b int) int {
		return a - b
	})

	if b.Size() != 0 {
		t.Errorf("size %d, want 0", b.Size())
	}

	if b.Height() != 0 {
		t.Errorf("height %d, want 0", b.Height())
	}

	if b.Degree() != 2 {
		t.Errorf("degree %d, want 2", b.Degree())
	}
}

func TestBTree_BTreeInsert_OneKey(t *testing.T) {
	b := New[int, int](2, func(a int, b int) int {
		return a - b
	})

	b.Insert(1, 1)

	inserts := []int{1}

	gotSize := b.Size()
	gotHeight := b.Height()

	if gotSize != 1 {
		t.Errorf("gotSize %d, wanted 1", gotSize)
	}

	if gotHeight != 0 {
		t.Errorf("gotHeight %d, wanted 0", gotHeight)
	}

	validateInsertsHelper(t, b, inserts)

}

func TestBTree_BTreeInsert_10Keys(t *testing.T) {
	b := New[int, int](2, func(a int, b int) int {
		return a - b
	})

	var inserts []int

	for i := 0; i < 10; i++ {
		inserts = append(inserts, i)
		b.Insert(i, i)
	}

	gotSize := b.Size()
	gotHeight := b.Height()

	if gotSize != 10 {
		t.Errorf("gotSize %d, wanted 10", gotSize)
	}

	if gotHeight != 2 {
		t.Errorf("gotHeight %d, wanted 2", gotHeight)
	}

	validateInsertsHelper(t, b, inserts)

}

func TestBTree_BTreeInsert_10KeysReverse(t *testing.T) {
	b := New[int, int](2, func(a int, b int) int {
		return a - b
	})

	var inserts []int

	for i := 9; i >= 0; i-- {
		inserts = append(inserts, 9-i)
		b.Insert(i, i)
	}

	gotSize := b.Size()
	gotHeight := b.Height()

	if gotSize != 10 {
		t.Errorf("gotSize %d, wanted 10", gotSize)
	}

	if gotHeight != 2 {
		t.Errorf("gotHeight %d, wanted 2", gotHeight)
	}

	validateInsertsHelper(t, b, inserts)

}

func TestBTree_InsertDuplicate(t *testing.T) {
	b := New[int, string](2, func(a int, b int) int {
		return a - b
	})

	ret, found := b.Insert(1, "first val")
	if found {
		t.Errorf("got non-nil value back from first insert. ret=%s", ret)
	}

	_, found = b.Insert(1, "second val")

	if !found {
		t.Errorf("got not found back for an overwrite. should have been \"first val\"")
	}
}

// degree two so insert 1 2 3 4 5 6 7 8 9 10 11 12 8
// second insert of 8 occurs when it is in a non-root, non-leaf node that is full
// it will be split, and elevated ot the parent. The implementation needs to check for this case
// TODO: add this to CS321 BTree tests
func TestBTree_InsertSplitMedianDuplicate(t *testing.T) {
	b := New[int, string](2, func(a int, b int) int {
		return a - b
	})

	ret, _ := b.Insert(1, "1")
	ret, _ = b.Insert(2, "2")
	ret, _ = b.Insert(3, "3")
	ret, _ = b.Insert(4, "4")
	ret, _ = b.Insert(5, "5")
	ret, _ = b.Insert(6, "6")
	ret, _ = b.Insert(7, "7")
	ret, _ = b.Insert(8, "8v1")
	ret, _ = b.Insert(9, "9")
	ret, _ = b.Insert(10, "10")
	ret, _ = b.Insert(11, "11")
	ret, _ = b.Insert(12, "12")
	// inserting a duplicate. the first K=8 is currently in a Full node that is not the root or a leaf
	ret, _ = b.Insert(8, "8v2")

	if ret != "8v1" {
		t.Errorf("missed the duplicate value during a split. should have been \"8v1\"")
	}

	if b.Size() != 12 {
		t.Errorf("still increased size, expected 12 but was %d", b.Size())
	}
}

// TODO add this to CS321 BTree tests
func TestBTree_InsertSplitNonLeafDuplicate(t *testing.T) {
	b := New[int, string](2, func(a int, b int) int {
		return a - b
	})

	ret, _ := b.Insert(1, "1")
	ret, _ = b.Insert(2, "2")
	ret, _ = b.Insert(3, "3")
	ret, _ = b.Insert(4, "4")
	ret, _ = b.Insert(5, "5")
	ret, _ = b.Insert(6, "6v1")
	ret, _ = b.Insert(7, "7")
	ret, _ = b.Insert(8, "8")
	ret, _ = b.Insert(9, "9")
	ret, _ = b.Insert(10, "10")
	ret, _ = b.Insert(11, "11")
	// inserting a duplicate. the first K=8 is currently in a Full node that is not the root or a leaf
	ret, _ = b.Insert(6, "6v2")

	if ret != "6v1" {
		t.Errorf("missed the duplicate value during a split. should have been \"6v1\"")
	}

	if b.Size() != 11 {
		t.Errorf("still increased size, expected 11 but was %d", b.Size())
	}
}

// TODO add this to CS321 BTree tests
func TestBTree_InsertSplitRootDuplicate(t *testing.T) {
	b := New[int, string](2, func(a int, b int) int {
		return a - b
	})

	ret, _ := b.Insert(1, "1")
	ret, _ = b.Insert(2, "2")
	ret, _ = b.Insert(3, "3")
	ret, _ = b.Insert(4, "4v1")
	ret, _ = b.Insert(5, "5")
	ret, _ = b.Insert(6, "6")
	ret, _ = b.Insert(7, "7")
	ret, _ = b.Insert(8, "8")
	ret, _ = b.Insert(8, "8")
	ret, _ = b.Insert(4, "4v2")

	if ret != "4v1" {
		t.Errorf("missed the duplicate value during a split. should have been \"4v1\"")
	}

	if b.Size() != 8 {
		t.Errorf("still increased size, expected 8 but was %d", b.Size())
	}
}

func validateInsertsHelper(t *testing.T, b *BTree[int, int], inserts []int) {

	var treeResult []int
	b.Traverse(func(v int) {
		treeResult = append(treeResult, v) // huh? slices are cool
	})

	for i := 0; i < b.Size(); i++ {
		if b.compare(treeResult[i], inserts[i]) != 0 {
			t.Errorf("Error inserts do not match. inserts[%d] = %d and tree value was %d", i, inserts[i], treeResult[i])
		}
	}
}

func valueSlice[K any, V any](b *BTree[K, V]) []V {
	var x []V
	b.Traverse(func(v V) {
		x = append(x, v)
	})
	return x
}
