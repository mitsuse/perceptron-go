package trie

type Trie interface {
	Get(key []int32) (Trie, bool)
	Update(key []int32) (Trie, bool)
}
