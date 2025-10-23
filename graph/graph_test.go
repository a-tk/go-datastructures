package graph

import "testing"

func Test_NewGraph(t *testing.T) {
	// goto function from aho-corasick paper
	g := NewGraphAutomata[int, string]()
	g.AddState(0)
	g.AddState(1)
	g.AddState(2)
	g.AddState(3)
	g.AddState(4)
	g.AddState(5)
	g.AddState(6)
	g.AddState(7)
	g.AddState(8)
	g.AddState(9)

	g.AddTransition(0, 1, "h")
	g.AddTransition(1, 2, "e")
	g.AddTransition(2, 8, "r")
	g.AddTransition(8, 9, "s")

	g.AddTransition(1, 6, "i")
	g.AddTransition(6, 7, "s")

	g.AddTransition(0, 3, "s")
	g.AddTransition(3, 4, "h")
	g.AddTransition(4, 5, "e")
}
