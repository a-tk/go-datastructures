package trie

// make will start with the smallest possible size before growing
// this should limit the total mem overhead from this change

type nodeMap struct {
	children map[rune]*nodeMap
	terminal bool
}

func newNodeMap(terminal bool) *nodeMap {
	return &nodeMap{
		make(map[rune]*nodeMap),
		terminal,
	}
}

type TrieMap struct {
	root *nodeMap
}

func NewTrieMap() *TrieMap {
	return &TrieMap{root: newNodeMap(false)}
}

func (t *TrieMap) AddWord(word string) {
	runes := []rune(word)
	x := t.root
	var y *nodeMap = nil
	for len(runes) > 0 {
		n := x.children[runes[0]]
		// if the path isn't there yet, add it
		if n == nil {
			x.children[runes[0]] = newNodeMap(len(runes) == 1)
		}
		y = x // remember the parent
		x = x.children[runes[0]]
		runes = runes[1:]
	}
	// out of runes, therefore a terminal
	if y != nil {
		y.terminal = true
	}
}

func (t *TrieMap) Search(word string) bool {
	runes := []rune(word)
	x := t.root
	for len(runes) > 0 {
		n := x.children[runes[0]]
		// if the path isn't there, we didn't find it
		if n == nil {
			return false
		}
		x = x.children[runes[0]]
		runes = runes[1:]
	}
	if x != nil {
		return x.terminal
	} else {
		return false
	}
}

//func (t *TrieMap) CountTerminals() int {
//	return t.root.countTermR()
//}
//
//func (x *nodeMap) countTermR() int {
//	if x != nil {
//		var i int
//		if x.terminal {
//			i = 1
//		} else {
//			i = 0
//		}
//		return i + x.child.countTermR() + x.sibling.countTermR()
//	} else {
//		return 0
//	}
//}

// TODO: autocomplete a partial n-gram
