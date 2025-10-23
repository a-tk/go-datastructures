package graph

// GraphAutomata is an adjacency list graph, in which states (vertices)
// are represented as the comparable type K, and edges are found according to
// a key type W
//
//	e.g. A -> B via 1 iff B appears in A's adjacency list AND is found by 1
type GraphAutomata[K, W comparable] struct {
	states map[K]map[W]K
}

func NewGraphAutomata[K, W comparable]() *GraphAutomata[K, W] {
	return &GraphAutomata[K, W]{
		states: make(map[K]map[W]K),
	}
}

func (g *GraphAutomata[K, W]) AddState(u K) (ok bool) {
	_, ok = g.states[u]
	// if u doesn't already exist
	if !ok {
		g.states[u] = make(map[W]K)
	}
	// if !ok, we created it -> true. If ok, it already existed -> false
	return !ok
}

func (g *GraphAutomata[K, W]) AddTransition(u, v K, w W) (ok bool) {
	m, ok := g.states[u]
	if ok {
		m[w] = v
	}
	return
}

func (g *GraphAutomata[K, W]) GetTransition(u K, w W) (v K, ok bool) {
	m, ok := g.states[u]
	if ok {
		v, ok = m[w]
	}
	return v, ok
}
