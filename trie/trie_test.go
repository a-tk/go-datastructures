package trie

import (
	"math/rand"
	"testing"
)

func genRandWord(n int) string {
	r := make([]rune, n)

	for i := 0; i < n; i++ {
		r[i] = rune(rand.Int() % 128)
	}
	return string(r)
}

func BenchmarkTrie_AddWord(b *testing.B) {
	t := New()

	for i := 0; i < b.N; i++ {
		t.AddWord(genRandWord(32))
	}
}

func BenchmarkTrie_AddWordI(b *testing.B) {
	t := New()

	for i := 0; i < b.N; i++ {
		t.AddWordI(genRandWord(32))
	}
}

func TestTrie_AddWord(t *testing.T) {
	trie := New()

	trie.AddWord("them")
	trie.AddWord("the")
	trie.AddWord("to")
	trie.AddWord("too")
	trie.AddWord("who")
	trie.AddWord("whose")
	trie.AddWord("世界")

	if trie.CountTerminals() != 7 {
		t.Errorf("error, not recording the correct number of terminals ")
	}
}

func TestTrie_SearchWord(t *testing.T) {
	trie := New()

	trie.AddWord("them")
	trie.AddWord("the")
	trie.AddWord("to")
	trie.AddWord("too")
	trie.AddWord("who")
	trie.AddWord("whose")

	if !trie.Search("them") {
		t.Errorf("didn't find them")
	}

	if !trie.Search("the") {
		t.Errorf("didn't find the")
	}

	if trie.Search("two") {
		t.Errorf("found two")
	}

	if trie.Search("whom") {
		t.Errorf("found whom")
	}

	if trie.Search("abcd") {
		t.Errorf("found abcd")
	}

	if trie.Search("") {
		t.Errorf("found empty string")
	}
}

func TestTrie_AddWordI(t *testing.T) {
	trie := New()

	trie.AddWord("them")
	trie.AddWord("the")
	trie.AddWord("to")
	trie.AddWord("too")
	trie.AddWord("who")
	trie.AddWord("whose")
}
