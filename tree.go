// Package freespacetree provides types designed to efficiently store a representation of free space on a storage device.
//
// Please see the repository at https://github.com/tomdionysus/freespacetree
package freespacetree

// Tree represents a tree of free space.
type Tree struct {
	root     *Node
	capacity uint64
}

// New returns a pointer to a new Tree with the specified capacity in blocks.
func New(capacity uint64) *Tree {
	inst := &Tree{
		root:     NewNode(0, capacity),
		capacity: capacity,
	}
	return inst
}

// Allocate attempts to allocate a continuous range of blocks in the Tree as specified, returning
// the first blockid in the allocation, and whether the space was successfully allocated.
func (tr *Tree) Allocate(blocks uint64) (uint64, bool) {
	return tr.root.Allocate(blocks)
}

// Deallocate frees the continuous space in the Tree referenced by the specified blockid and length.
func (tr *Tree) Deallocate(blockid uint64, blocklength uint64) {
	tr.root = tr.root.Deallocate(blockid, blocklength)
}
