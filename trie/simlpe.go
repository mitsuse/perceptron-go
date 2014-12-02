package trie

import (
	"sort"
)

type SimpleTrie struct {
	char     int32
	value    int
	childSeq SimpleTrieSeq
}

func NewSimpleTrie() *SimpleTrie {
	t := &SimpleTrie{
		childSeq: make(SimpleTrieSeq, 0),
	}

	return t
}

func (t *SimpleTrie) Get(key []int32) (Trie, bool) {
	node, read := t.find(key)

	return node, len(key) == read
}

func (t *SimpleTrie) Update(key []int32) (Trie, bool) {
	node, read := t.find(key)
	found := len(key) == read

	for ; read < len(key); read++ {
		child := &SimpleTrie{
			char:     key[read],
			value:    0,
			childSeq: make(SimpleTrieSeq, 0),
		}

		node.childSeq = append(node.childSeq, child)
		sort.Sort(node.childSeq)

		node = child
	}

	return node, found
}

func (t *SimpleTrie) find(key []int32) (node *SimpleTrie, read int) {
	node = t

	for keyOffset, char := range key {
		if len(node.childSeq) == 0 {
			return node, keyOffset
		}

		childOffset := sort.Search(len(node.childSeq)-1, func(offset int) bool {
			return node.childSeq[offset].char >= char
		})
		child := node.childSeq[childOffset]

		if child.char != char {
			return node, keyOffset
		}

		node = child
	}

	return node, len(key)
}

type SimpleTrieSeq []*SimpleTrie

func (s SimpleTrieSeq) Len() int {
	return len(s)
}

func (s SimpleTrieSeq) Less(i, j int) bool {
	return s[i].char < s[j].char
}

func (s SimpleTrieSeq) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
