package kdtree

type node[V any] struct {
	val  V
	l, r *node[V]
	disc int
}

// V is an object that contains k keys, according to k dimensions
type KDTree[V any] struct {
	k       int
	root    *node[V]
	compare func(V, V, int) int
}

func newNode[V any](v V, d int) *node[V] {
	return &node[V]{
		val:  v,
		disc: d,
	}
}

// any comparison needs to know which dimension to compare against
func New[V any](compare func(V, V, int) int, k int) (t *KDTree[V], ok bool) {
	if k < 0 {
		return t, false
	} else {
		return &KDTree[V]{k: k, root: nil, compare: compare}, true
	}
}

//func (t *BST[K, V]) Insert(k K, v V) (oldValue V, replaced bool) {
//
//	x := t.root
//	var y *node[K, V] = nil
//	z := newNode(k, v)
//	for x != nil {
//		y = x
//		cmp := t.compare(z.key, x.key)
//		if cmp == 0 {
//			prev := x.val
//			x.val = v
//			return prev, true
//		} else if cmp < 0 {
//			x = x.l
//		} else {
//			x = x.r
//		}
//	}
//	z.p = y
//	if y == nil {
//		t.root = z
//	} else if t.compare(z.key, y.key) < 0 {
//		y.l = z
//	} else {
//		y.r = z
//	}
//	return oldValue, false
//}

// the normal BST insert algorithm should work fine here, just deal with the discriminator
func (t *KDTree[V]) Insert(v V) (dup bool) {
	x := t.root
	var y *node[V] = nil
	z := newNode(v, 0)
	for x != nil {
		y = x
		cmp := t.compare(z.val, x.val, x.disc)
		if cmp == 0 {
			// TODO: compare other fields for equality with for loop
			return true
		} else if cmp < 0 {
			x = x.l
		} else {
			x = x.r
		}
	}
	if y == nil {
		t.root = z
		t.root.disc = 0
	} else if t.compare(z.val, y.val, y.disc) < 0 {
		y.l = z
		z.disc = y.disc + 1%t.k
	} else {
		y.r = z
		z.disc = y.disc + 1%t.k
	}
	return false
}
