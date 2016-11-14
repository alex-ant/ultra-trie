package trie

// node represents a separate tree node.
type node struct {
	key      string
	children map[string]node
	data     interface{}
}

// Tree contains Trie structure.
type Tree struct {
	children map[string]node
}

// New returns new Tree.
func New() *Tree {
	return &Tree{
		children: make(map[string]node),
	}
}
