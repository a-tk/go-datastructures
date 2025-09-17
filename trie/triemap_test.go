package trie

import "testing"

func TestTrieMap_AddWord(t *testing.T) {
	m := NewTrieMap()

	m.AddWord("he")
	m.AddWord("hell")

}

func TestTrieMap_SearchWord(t *testing.T) {
	trie := NewTrieMap()

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
