package trie

type Trie interface {
	Get(key []int32) (Trie, error)
	Update(key []int32) (Trie, error)
}
