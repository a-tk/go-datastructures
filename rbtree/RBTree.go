package rbtree

const (
	RED   = 0
	BLACK = 1
)

type node[K any, V any] struct {
	key     K
	val     V
	l, r, p *node[K, V]
	color   int // consider smaller data size
}

type RBTree[K any, V any] struct {
	root    *node[K, V]
	compare func(K, K) int
	NIL     *node[K, V]
}

func (t *RBTree[K, V]) newNode(k K, v V) *node[K, V] {
	return &node[K, V]{
		key:   k,
		val:   v,
		color: RED,
		l:     t.NIL,
		r:     t.NIL,
		p:     t.NIL,
	}
}

func New[K any, V any](compare func(K, K) int) *RBTree[K, V] {
	NIL := &node[K, V]{
		l:     nil,
		r:     nil,
		p:     nil,
		color: BLACK,
	}
	return &RBTree[K, V]{
		NIL:     NIL,
		root:    NIL,
		compare: compare}
}

func (t *RBTree[K, V]) Insert(k K, v V) (old V, replaced bool) {
	x := t.root
	y := t.NIL
	z := t.newNode(k, v) // all new nodes are RED
	for x != t.NIL {
		y = x
		cmp := t.compare(z.key, x.key)
		if cmp == 0 {
			prev := x.val
			x.val = v
			//tree was not mutated, no fixing needed
			return prev, true
		} else if cmp < 0 {
			x = x.l
		} else {
			x = x.r
		}
	}
	z.p = y
	if y == t.NIL {
		t.root = z
	} else if t.compare(z.key, y.key) < 0 {
		y.l = z
	} else {
		y.r = z
	}
	t.insertFixup(z)
	return old, false
}
func (t *RBTree[K, V]) insertFixup(z *node[K, V]) {
	for z.p.color == RED {
		if z.p == z.p.p.l {
			y := z.p.p.r //z's uncle
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.r {
					z = z.p
					t.leftRotate(z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				t.rightRotate(z.p.p)
			}
		} else {
			y := z.p.p.l
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.l {
					z = z.p
					t.rightRotate(z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				t.leftRotate(z.p.p)
			}
		}
	}
	t.root.color = BLACK
}

func (t *RBTree[K, V]) leftRotate(x *node[K, V]) {
	y := x.r
	x.r = y.l
	if y.l != t.NIL {
		y.l.p = x
	}
	y.p = x.p
	if x.p == t.NIL {
		t.root = y
	} else if x == x.p.l {
		x.p.l = y
	} else {
		x.p.r = y
	}
	y.l = x
	x.p = y
}

func (t *RBTree[K, V]) rightRotate(x *node[K, V]) {
	y := x.l
	x.l = y.r
	if y.r != t.NIL {
		y.r.p = x
	}
	y.p = x.p
	if x.p == t.NIL {
		t.root = y
	} else if x == x.p.r {
		x.p.r = y
	} else {
		x.p.l = y
	}
	y.r = x
	x.p = y
}

// Height calculates how many nodes from the top of the tree
// to the bottom of the tree on the longest path, including the root
// this means that the number of possible nodes at each height does not
// obey the theoretical rule of n=2^h (i.e., height 1 (the root) has one node, not two)
func (t *RBTree[K, V]) Height() int {
	return t.height(t.root)
}

func (t *RBTree[K, V]) height(x *node[K, V]) int {
	if x == t.NIL {
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

func (t *RBTree[K, V]) rbtransplant(u *node[K, V], v *node[K, V]) {
	if u.p == t.NIL {
		t.root = v
	} else if u == u.p.l {
		u.p.l = v
	} else {
		u.p.r = v
	}
	v.p = u.p
}

func (t *RBTree[K, V]) Remove(k K) (old V, found bool) {
	// search to get the node, then remove it
	x := t.search(t.root, k)
	if x != t.NIL {
		t.remove(x)
		return x.val, true
	}
	return old, false
}

func (t *RBTree[K, V]) remove(z *node[K, V]) {
	y := z
	var x *node[K, V]
	y_original_color := y.color
	if z.l == t.NIL {
		t.rbtransplant(z, z.r)
	} else if z.r == t.NIL {
		x = z.l
		t.rbtransplant(z, z.l)
	} else {
		y := t.minimum(z.r)
		y_original_color = y.color
		x = y.r
		if y != z.r {
			t.rbtransplant(y, y.r)
			y.r = z.r
			y.r.p = y
		} else {
			x.p = y
		}
		t.rbtransplant(z, y)
		y.l = z.l
		y.l.p = y
		y.color = z.color
	}
	if y_original_color == BLACK {
		t.removeFixup(x)
	}
}
func (t *RBTree[K, V]) removeFixup(x *node[K, V]) {
	for x != t.root && x.color == BLACK {
		if x == x.p.l {
			w := x.p.r
			if w.color == RED {
				w.color = BLACK
				x.p.color = RED
				t.leftRotate(x.p)
				w = x.p.r
			}
			if w.l.color == BLACK && w.r.color == BLACK {
				w.color = RED
				x = x.p
			} else {
				if w.r.color == BLACK {
					w.l.color = BLACK
					w.color = RED
					t.rightRotate(w)
					w = x.p.r
				}
				w.color = x.p.color
				x.p.color = BLACK
				w.r.color = BLACK
				t.leftRotate(x.p)
				x = t.root
			}
		} else {
			w := x.p.l
			if w.color == RED {
				w.color = BLACK
				x.p.color = RED
				t.rightRotate(x.p)
				w = x.p.l
			}
			if w.r.color == BLACK && w.l.color == BLACK {
				w.color = RED
				x = x.p
			} else {
				if w.l.color == BLACK {
					w.r.color = BLACK
					w.color = RED
					t.leftRotate(w)
					w = x.p.l
				}
				w.color = x.p.color
				x.p.color = BLACK
				w.l.color = BLACK
				t.rightRotate(x.p)
				x = t.root
			}
		}
	}
	x.color = BLACK
}

func (t *RBTree[K, V]) Search(k K) (val V, found bool) {
	x := t.search(t.root, k)
	if x == t.NIL {
		return val, false
	} else {
		return x.val, true
	}
}

func (t *RBTree[K, V]) search(x *node[K, V], k K) *node[K, V] {

	for x != t.NIL && t.compare(x.key, k) != 0 {
		if t.compare(k, x.key) < 0 {
			x = x.l
		} else {
			x = x.r
		}
	}
	return x
}

func (t *RBTree[K, V]) Clear() {
	// two thoughts. If t.root is set to nil, the tree gets garbage collected?
	// alternatively, traverse the tree removing each node recursively <- safer
	t.traverse(t.root, func(n *node[K, V]) {
		t.remove(n)
	})
}

func (t *RBTree[K, V]) Size() int {
	return t.size(t.root)
}

func (t *RBTree[K, V]) size(x *node[K, V]) int {
	if x == t.NIL {
		return 0
	} else {
		return 1 + t.size(x.l) + t.size(x.r)
	}
}

func (t *RBTree[K, V]) ContainsKey(k K) bool {
	x := t.search(t.root, k)
	if x != t.NIL {
		return true
	} else {
		return false
	}
}

func (t *RBTree[K, V]) ContainsValue(v V, cmp func(V, V) int) bool {
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

func (t *RBTree[K, V]) minimum(x *node[K, V]) *node[K, V] {
	for x.l != t.NIL {
		x = x.l
	}
	return x
}

func (t *RBTree[K, V]) maximum(x *node[K, V]) *node[K, V] {
	for x.r != t.NIL {
		x = x.r
	}
	return x
}

func (t *RBTree[K, V]) Successor(k K) (val V, found bool) {
	x := t.search(t.root, k)
	x = t.successor(x)
	if x != t.NIL {
		return x.val, true
	} else {
		return val, false
	}
}

func (t *RBTree[K, V]) successor(x *node[K, V]) *node[K, V] {
	if x.r != t.NIL {
		return t.minimum(x.r)
	} else {
		y := x.p
		for y != t.NIL && x == y.r {
			x = y
			y = y.p
		}
		return y
	}
}

func (t *RBTree[K, V]) Predecessor(k K) (val V, found bool) {
	x := t.search(t.root, k)
	x = t.predecessor(x)
	if x != t.NIL {
		return x.val, true
	} else {
		return val, false
	}
}

func (t *RBTree[K, V]) predecessor(x *node[K, V]) *node[K, V] {
	if x.l != t.NIL {
		return t.maximum(x.l)
	} else {
		y := x.p
		for y != t.NIL && x == y.l {
			x = y
			y = y.p
		}
		return y
	}
}

func (t *RBTree[K, V]) traverse(x *node[K, V], action func(*node[K, V])) {
	if x != t.NIL {
		t.traverse(x.l, action)
		action(x)
		t.traverse(x.r, action)
	}
}
