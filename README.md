# Go Data Structures Library

* Author: a-tk, Andre Keys

## Overview

This project is a collection of common data structure implementations in Go.
It is not intended to be a collections framework or library, but rather provide 
idiomatic, easy-to-understand implementations that demonstrate 
core algorithms and data organization techniques.

Included structures:

- **Binary Search Tree (BST)**
- **B-Tree (in-memory only, configurable degree)**
- **Gap Buffer**
- **Deque**
- **Heap**
- **k-d Tree**
- **Circular Linked List**
- **Single Linked List**
- **Double Linked List**
- **LRU Cache**
- **Queue**
- **Red-Black Tree**
- **Stack**
- **Trie**

## Compiling and Using

To build the project:

```bash
go build ./...

To run the included tests:

```
go test ./...

### Usage

import "datastructures/bst"

cmp := func(a, b int) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    }
    return 0
}

tree := bst.New[int, string](cmp)
tree.Insert(5, "five")
tree.Insert(2, "two")
tree.Insert(7, "seven")

val, ok := tree.Search(2)
// val == "two", ok == true


import "datastructures/trie"

t := trie.New()
t.Insert("hello")
t.Insert("hell")
t.Insert("heaven")

found := t.Search("hello") // true
miss  := t.Search("hero")  // false


Other structures can be imported similarly by their package path
(e.g., datastructures/deque, datastructures/heap, datastructures/stack, etc.).

No direct command-line interface is provided.

