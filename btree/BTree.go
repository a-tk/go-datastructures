// a write to disk BTree
package btree

// TODO: consider the encoding package: https://pkg.go.dev/encoding@go1.24.5
// make the node impl the interface? then be in charge of where to write it?
// Ask for the length of the byte slice, to figure out where/how to write to the disk maybe?

// here, an array of pointers offers more advantages than direct object storage in an array
// for sparse trees, 50% space in the arrays may be wasted
// this consideration is only for trees held entirely in memory
type container[K any, V any] struct {
	key    K
	valptr *V
}

// TODO: delete :(
type node[K any, V any] struct {
	n        int
	leaf     bool
	keys     []container[K, V]
	children []*node[K, V]
}

type BTree[K any, V any] struct {
	root    *node[K, V]
	degree  int
	height  int
	size    int
	compare func(K, K) int
}

func newNode[K any, V any](t int) *node[K, V] {
	return &node[K, V]{
		n:        0,
		leaf:     true,
		keys:     make([]container[K, V], 2*t-1),
		children: make([]*node[K, V], 2*t),
	}
}

func NewBTree[K any, V any](degree int, compare func(K, K) int) (b *BTree[K, V]) {
	return &BTree[K, V]{
		degree:  degree,
		root:    newNode[K, V](degree),
		height:  0,
		size:    0,
		compare: compare,
	}
}

func (b *BTree[K, V]) Search(k K) *V {
	return b.search(b.root, k)
}

func (b *BTree[K, V]) search(x *node[K, V], k K) *V {
	//can be updated to use the bsearch impl below
	i := 0
	for i < x.n && b.compare(x.keys[i].key, k) < 1 {
		i++
	}
	if i < x.n && b.compare(x.keys[i].key, k) == 0 {
		return x.keys[i].valptr
	} else if x.leaf {
		return nil
	} else {
		return b.search(x.children[i], k)
	}
}

func (b *BTree[K, V]) Insert(k K, v *V) *V {
	r := b.root
	if r.n == 2*b.degree-1 {
		s := b.splitRoot()
		return b.insertNonFull(s, k, v)
	} else {
		return b.insertNonFull(r, k, v)
	}
}

func (b *BTree[K, V]) insertNonFull(x *node[K, V], k K, v *V) *V {

	// new strategy: search the nodes list before proceeding. If there is a duplicate, deal with it
	// if not proceed per pseudocode, but beware of the case were splitting a child elevates a duplicate
	// for degree 2000 and linear search, time spent was 2x
	// iterative bsearch results in a few additional ns/op

	i := b.iterativeBSearch(x.keys, k, x.n)
	if i != -1 {
		//duplicate
		previous := x.keys[i].valptr
		x.keys[i] = container[K, V]{key: k, valptr: v}
		return previous
	} else if x.leaf {
		//insert into a leaf
		i = x.n - 1
		for i >= 0 && b.compare(x.keys[i].key, k) > 0 {
			x.keys[i+1] = x.keys[i]
			i = i - 1
		}

		x.keys[i+1] = container[K, V]{key: k, valptr: v}
		b.size++
		x.n = x.n + 1
		return nil
	} else {
		//search for the correct child to continue looking
		i = x.n - 1
		for i >= 0 && b.compare(x.keys[i].key, k) > 0 {
			i = i - 1
		}
		i = i + 1

		if x.children[i].n == 2*b.degree-1 {
			b.splitChild(x, i)
			// which child to insert into?
			if b.compare(x.keys[i].key, k) < 0 {
				i = i + 1
			} else if b.compare(x.keys[i].key, k) == 0 { // check to see if the median value is the same K we are inserting
				previous := x.keys[i].valptr                   // save previous value
				x.keys[i] = container[K, V]{key: k, valptr: v} // replace it with the new value
				return previous
			}
		}
		return b.insertNonFull(x.children[i], k, v)
	}
}

func (b *BTree[K, V]) splitChild(x *node[K, V], i int) {
	t := b.degree
	y := x.children[i]
	z := newNode[K, V](t)
	z.leaf = y.leaf
	z.n = t - 1
	for j := 0; j <= t-2; j++ {
		z.keys[j] = y.keys[j+t]
	}
	if !y.leaf {
		for j := 0; j <= t-1; j++ {
			z.children[j] = y.children[j+t]
		}
	}
	y.n = t - 1
	for j := x.n; j >= i+1; j-- {
		x.children[j+1] = x.children[j]
	}
	x.children[i+1] = z
	for j := x.n - 1; j >= i; j-- {
		x.keys[j+1] = x.keys[j]
	}
	x.keys[i] = y.keys[t-1]
	x.n = x.n + 1

	// old values still might exist in the nodes from splitting, even though they are not considered
	// valid by the tree, which would alias valid pointers. Even if delete is implemented, there would
	//remain aliased references
	// nil out the keys and children of y (who had the top t values removed)
	// note the relative position of the keys and the children (draw a picture is helpful)
	for i := y.n; i < 2*b.degree-1; i++ {
		y.keys[i].valptr = nil
		y.children[i+1] = nil
	}
}

func (b *BTree[K, V]) splitRoot() *node[K, V] {
	s := newNode[K, V](b.degree)
	s.leaf = false
	s.n = 0
	s.children[0] = b.root
	b.root = s
	b.splitChild(s, 0)
	b.height++
	return s
}

func (b *BTree[K, V]) Height() int {
	return b.height
}

func (b *BTree[K, V]) Size() int {
	return b.size
}

func (b *BTree[K, V]) Degree() int {
	return b.degree
}

func (b *BTree[K, V]) Traverse(action func(*V)) {
	b.traverse(b.root, action)
}

func (b *BTree[K, V]) traverse(x *node[K, V], action func(*V)) {
	if x == nil {
		return
	} else if x.leaf {
		for i := 0; i < x.n; i++ {
			action(x.keys[i].valptr)
		}
	} else {
		for i := 0; i < x.n; i++ {
			b.traverse(x.children[i], action)
			action(x.keys[i].valptr)
		}
		b.traverse(x.children[x.n], action)
	}
}

func (b *BTree[K, V]) iterativeBSearch(keys []container[K, V], k K, n int) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := low + (high-low)/2 // Prevents potential overflow compared to (low + high) / 2

		if b.compare(keys[mid].key, k) == 0 {
			return mid
		} else if b.compare(keys[mid].key, k) < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
