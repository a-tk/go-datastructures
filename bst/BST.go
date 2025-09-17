package bst

import "datastructures/stack"

type node[K any, V any] struct {
	key     K
	val     V
	l, r, p *node[K, V]
}

type BST[K any, V any] struct {
	root    *node[K, V]
	compare func(K, K) int
}

func newNode[K any, V any](k K, v V) *node[K, V] {
	return &node[K, V]{
		key: k,
		val: v,
	}
}

func New[K any, V any](compare func(K, K) int) *BST[K, V] {
	return &BST[K, V]{root: nil, compare: compare}
}

func (t *BST[K, V]) Insert(k K, v V) (oldValue V, replaced bool) {

	x := t.root
	var y *node[K, V] = nil
	z := newNode(k, v)
	for x != nil {
		y = x
		cmp := t.compare(z.key, x.key)
		if cmp == 0 {
			prev := x.val
			x.val = v
			return prev, true
		} else if cmp < 0 {
			x = x.l
		} else {
			x = x.r
		}
	}
	z.p = y
	if y == nil {
		t.root = z
	} else if t.compare(z.key, y.key) < 0 {
		y.l = z
	} else {
		y.r = z
	}
	return oldValue, false
}

// Height calculates how many nodes from the top of the tree
// to the bottom of the tree on the longest path, including the root
// this means that the number of possible nodes at each height does not
// obey the theoretical rule of n=2^h (i.e., height 1 (the root) has one node, not two)
func (t *BST[K, V]) Height() int {
	return t.height(t.root)
}

func (t *BST[K, V]) height(x *node[K, V]) int {
	if x == nil {
		return 0
	} else {
		l := t.height(x.l)
		r := t.height(x.r)
		if l < r {
			return r + 1
		} else {
			return l + 1
		}
	}
}

func (t *BST[K, V]) transplant(u *node[K, V], v *node[K, V]) {
	if u.p == nil {
		t.root = v
	} else if u == u.p.l {
		u.p.l = v
	} else {
		u.p.r = v
	}
	if v != nil {
		v.p = u.p
	}
}

func (t *BST[K, V]) Remove(k K) (oldValue V, found bool) {
	// search to get the node, then remove it
	x := t.search(t.root, k)
	if x != nil {
		t.remove(x)
		return x.val, true
	}
	return oldValue, false
}

func (t *BST[K, V]) remove(z *node[K, V]) {
	if z.l == nil {
		t.transplant(z, z.r)
	} else if z.r == nil {
		t.transplant(z, z.l)
	} else {
		y := t.minimum(z.r)
		if y != z.r {
			t.transplant(y, y.r)
			y.r = z.r
			y.r.p = y
		}
		t.transplant(z, y)
		y.l = z.l
		y.l.p = y
	}
}

func (t *BST[K, V]) Search(k K) (val V, found bool) {
	x := t.search(t.root, k)
	if x == nil {
		return val, false
	} else {
		return x.val, true
	}
}

func (t *BST[K, V]) search(x *node[K, V], k K) *node[K, V] {

	for x != nil && t.compare(x.key, k) != 0 {
		if t.compare(k, x.key) < 0 {
			x = x.l
		} else {
			x = x.r
		}
	}
	return x
}

func (t *BST[K, V]) Clear() {
	// two thoughts. If t.root is set to nil, the tree gets garbage collected?
	// alternatively, traverse the tree removing each node recursively <- safer
	t.traverse(t.root, func(n *node[K, V]) {
		t.remove(n)
	})
}

func (t *BST[K, V]) Size() int {
	return t.size(t.root)
}

func (t *BST[K, V]) size(x *node[K, V]) int {
	if x == nil {
		return 0
	} else {
		return 1 + t.size(x.l) + t.size(x.r)
	}
}

func (t *BST[K, V]) ContainsKey(k K) bool {
	x := t.search(t.root, k)
	if x != nil {
		return true
	} else {
		return false
	}
}

func (t *BST[K, V]) ContainsValue(v V, cmp func(V, V) int) bool {
	found := false

	// can't stop the traversal :(
	t.traverse(t.root, func(n *node[K, V]) {
		if cmp(n.val, v) == 0 {
			found = true
		}
	})
	// java TreeMap finds the minimum then each successor, clever!
	return found
}

func (t *BST[K, V]) minimum(x *node[K, V]) *node[K, V] {
	for x.l != nil {
		x = x.l
	}
	return x
}

func (t *BST[K, V]) maximum(x *node[K, V]) *node[K, V] {
	for x.r != nil {
		x = x.r
	}
	return x
}

func (t *BST[K, V]) Successor(k K) (val V, found bool) {
	x := t.search(t.root, k)
	x = t.successor(x)
	if x != nil {
		return x.val, true
	} else {
		return val, false
	}
}

func (t *BST[K, V]) successor(x *node[K, V]) *node[K, V] {
	if x.r != nil {
		return t.minimum(x.r)
	} else {
		y := x.p
		for y != nil && x == y.r {
			x = y
			y = y.p
		}
		return y
	}
}

func (t *BST[K, V]) Predecessor(k K) (val V, found bool) {
	x := t.search(t.root, k)
	x = t.predecessor(x)
	if x != nil {
		return x.val, true
	} else {
		return val, false
	}
}

func (t *BST[K, V]) predecessor(x *node[K, V]) *node[K, V] {
	if x.l != nil {
		return t.maximum(x.l)
	} else {
		y := x.p
		for y != nil && x == y.l {
			x = y
			y = y.p
		}
		return y
	}
}

func (t *BST[K, V]) traverse(x *node[K, V], action func(*node[K, V])) {
	if x != nil {
		t.traverse(x.l, action)
		action(x)
		t.traverse(x.r, action)
	}
}

// from CLRS
// Hint: an easy solution uses a stack as an auxiliary data structure.
// A more complicated, but elegant, solution uses no stack but assumes
// that we can test two pointers for equality.
func (t *BST[K, V]) traverseIter(action func(*node[K, V])) {
	s := stack.New[*node[K, V]]()

	current := t.root
	// two conditions. On the first round, nothing in the queue but current is not nil
	// traversal to the top of the tree also leaves the stack empty, but current not nil
	// at the leaves, current is nil but stack has elements in it.
	for !s.Empty() || current != nil {
		for current != nil {
			s.Push(current)
			current = current.l
		}

		current, _ = s.Pop()
		action(current)
		current = current.r
	}
}
