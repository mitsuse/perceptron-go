package perceptron

import (
	"github.com/mitsuse/perceptron-go/trie"
)

type Indexer struct {
	trie *trie.SimpleTrie
	size int
}

func NewIndexer() *Indexer {
	i := &Indexer{
		trie: trie.NewSimpleTrie(),
		size: 0,
	}

	return i
}

func (i *Indexer) Size() int {
	return i.size
}

func (i *Indexer) Index(identifier []int32, indexed bool) int {
	var index int

	if indexed {
		node, exist := i.trie.Update(identifier)
		if !exist {
			i.size++
			node.SetValue(i.size)
		}

		index = node.Value()
	} else {
		node, exist := i.trie.Get(identifier)
		if exist {
			index = node.Value()
		}
	}

	return index
}
