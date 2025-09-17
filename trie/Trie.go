package trie

// uses left-child right sibling representation

type node struct {
	r              rune
	child, sibling *node
	terminal       bool
}

type Trie struct {
	root *node
}

func New() *Trie {
	return &Trie{}
}

func (t *Trie) AddWordI(word string) {
	runes := []rune(word)
	x := t.root
	var y *node = nil
	for len(runes) > 0 {

		// if no node exists on the path, add one
		if x == nil {
			x = &node{
				r: runes[0],
			}
			// link it to the parent
			if y != nil {
				y.child = x
			}
		} else if x.r == runes[0] {
			y = x
			x = x.child

		} else { // keep searching
			y = x
			x = x.sibling
		}
		runes = runes[1:]
	}
	if y != nil {
		y.terminal = true
	}
}

func (t *Trie) AddWord(word string) {
	t.root = t.root.addWordR([]rune(word))
}

func (x *node) addWordR(runes []rune) *node {
	// base case
	if len(runes) == 0 {
		return nil
	}

	if x == nil { // no node exists yet in the search path
		x = &node{r: runes[0]}
		x.terminal = len(runes) == 1 // if this is the last character, the node is a terminal
		x.child = x.child.addWordR(runes[1:])
	} else {

		// if the rune matches
		if x.r == runes[0] {
			// and if it's the last one, set terminal
			if len(runes) == 1 {
				x.terminal = true
			} else {
				// if it wasn't the last one, keep searching/adding
				//  note, if this is not protected against the len(runes) == 1 condition
				//  t.addWordR would return nil, and break the tree links
				x.child = x.child.addWordR(runes[1:])
			}
		} else {
			// didn't match, check the sibling
			x.sibling = x.sibling.addWordR(runes)
		}
	}
	return x
}

func (t *Trie) Search(word string) bool {
	return t.root.searchR([]rune(word))
}

func (x *node) searchR(runes []rune) bool {
	if len(runes) == 0 {
		return false
	}
	if x == nil {
		return false
	} else {
		// if it's the last character, return x.terminal
		if x.r == runes[0] {
			if len(runes) == 1 {
				return x.terminal
			} else {
				return x.child.searchR(runes[1:])
			}
		} else {
			return x.sibling.searchR(runes)
		}
	}
}

func (t *Trie) CountTerminals() int {
	return t.root.countTermR()
}

func (x *node) countTermR() int {
	if x != nil {
		var i int
		if x.terminal {
			i = 1
		} else {
			i = 0
		}
		return i + x.child.countTermR() + x.sibling.countTermR()
	} else {
		return 0
	}
}

// TODO: autocomplete a partial n-gram
