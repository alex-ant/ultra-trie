package trie

import "sync"

// node represents a separate tree node.
type node struct {
	key      byte
	children map[byte]*node
	data     interface{}
}

func newEmptyNode(key byte) *node {
	return &node{
		key:      key,
		children: make(map[byte]*node),
	}
}

func (n *node) createPathChildren(path string, data interface{}) {
	key := path[0]

	_, childExists := n.children[key]
	if !childExists {
		n.children[key] = newEmptyNode(key)
	}

	cutPath := path[1:len(path)]
	if len(cutPath) == 0 {
		n.children[key].data = data
		return
	}
	n.children[key].createPathChildren(cutPath, data)
}

func (n *node) lookupPathChildren(path string) bool {
	key := path[0]

	_, childExists := n.children[key]
	if !childExists {
		return false
	}

	cutPath := path[1:len(path)]
	if len(cutPath) == 0 {
		return true
	}

	return n.children[key].lookupPathChildren(cutPath)
}

////////////////////////////////

// Tree contains Trie structure.
type Tree struct {
	mu       sync.Mutex
	children map[byte]*node
}

// New returns new Tree.
func New() *Tree {
	return &Tree{
		children: make(map[byte]*node),
	}
}

// PrefixExists tells whether the requested prefix exists in the tree.
func (t *Tree) PrefixExists(path string) bool {
	// skip if no path provided
	if path == "" {
		return false
	}

	// manage mutex
	t.mu.Lock()
	defer t.mu.Unlock()

	// check the base node for prefix
	key := path[0]
	_, baseNodeExists := t.children[key]
	if !baseNodeExists {
		return false
	}

	// return true if the node is the last one requested
	if len(path) == 1 {
		return true
	}

	// lookup path children
	return t.children[key].lookupPathChildren(path[1:len(path)])
}

// Add adds a new record to the tree.
func (t *Tree) Add(path string, data interface{}) {
	// skip if no path provided
	if path == "" {
		return
	}

	// manage mutex
	t.mu.Lock()
	defer t.mu.Unlock()

	// make base node if doesn't exist
	key := path[0]
	_, baseNodeExists := t.children[key]
	if !baseNodeExists {
		t.children[key] = newEmptyNode(key)
	}

	// if the base node is the last one requested, adding the data right into it
	if len(path) == 1 {
		t.children[key].data = data
		return
	}

	// creating chain children
	t.children[key].createPathChildren(path[1:len(path)], data)
}
