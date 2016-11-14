package trie

import (
	"fmt"
	"sync"
)

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

// GetPrefixMembers returns a slice of prefix members. Error if the prefix doesn't exist.
func (t *Tree) GetPrefixMembers(path string) (mm []Member, err error) {
	// check whether the prefix exists
	if !t.PrefixExists(path) {
		err = fmt.Errorf("prefix %s doesn't exist", path)
		return
	}

	// manage mutex
	t.mu.Lock()
	defer t.mu.Unlock()

	// get prefix node

	// get base node first
	key := path[0]
	baseNode, baseNodeExists := t.children[key]
	if !baseNodeExists {
		err = fmt.Errorf("base node %s doesn't exist", string(key))
		return
	}

	// return data if the node is the last one requested
	if len(path) == 1 {
		mm = baseNode.getAllMembers()
		return
	}

	// going deeper
	var n *node
	n, err = baseNode.getChildNodeByPath(path[1:len(path)])
	if err != nil {
		return
	}

	mm = n.getAllMembers()
	return
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
		t.children[key] = newEmptyNode(string(key))
	}

	// if the base node is the last one requested, adding the data right into it
	if len(path) == 1 {
		t.children[key].data = data
		return
	}

	// creating chain children
	t.children[key].createPathChildren(string(key), path[1:len(path)], data)
}
