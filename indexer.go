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
	// TODO: Implement this.
	return 0
}

func (i *Indexer) Index(identifier []int32, indexed bool) int {
	// TODO: Implement this.
	return 0
}
