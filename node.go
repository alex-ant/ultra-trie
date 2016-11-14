package trie

import "fmt"

// node represents a separate tree node.
type node struct {
	key      string
	children map[byte]*node
	data     interface{}
}

func newEmptyNode(key string) *node {
	return &node{
		key:      key,
		children: make(map[byte]*node),
	}
}

// Member contains prefix's member data.
type Member struct {
	Key  string
	Data interface{}
}

func (n *node) getAllMembers() (mm []Member) {
	// append node's data
	if n.data != nil {
		mm = append(mm, Member{
			Key:  n.key,
			Data: n.data,
		})
	}

	// loop through children
	for _, v := range n.children {
		mm = append(mm, v.getAllMembers()...)
	}

	return
}

func (n *node) createPathChildren(parentKey, path string, data interface{}) {
	key := path[0]

	currentKey := parentKey + string(key)

	_, childExists := n.children[key]
	if !childExists {
		n.children[key] = newEmptyNode(currentKey)
	}

	cutPath := path[1:len(path)]
	if len(cutPath) == 0 {
		n.children[key].data = data
		return
	}
	n.children[key].createPathChildren(currentKey, cutPath, data)
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

func (n *node) getChildNodeByPath(path string) (*node, error) {
	key := path[0]

	childNode, childExists := n.children[key]
	if !childExists {
		return nil, fmt.Errorf("child by the key %s doesn't exist", string(key))
	}

	cutPath := path[1:len(path)]
	if len(cutPath) == 0 {
		return childNode, nil
	}

	return n.children[key].getChildNodeByPath(cutPath)
}
