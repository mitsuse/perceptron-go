package trie

type Trie interface {
	Value() int
	SetValue(value int)
	Get(key []int32) (Trie, bool)
	Update(key []int32) (Trie, bool)
}
