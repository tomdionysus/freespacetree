package freespacetree

// Node represents a range of free space in the tree.
type Node struct {
	from  uint64
	to    uint64
	left  *Node
	right *Node
}

// NewNode returns a pointer to a new Node with the specified range of free space.
func NewNode(from, to uint64) *Node {
	inst := &Node{
		left:  nil,
		right: nil,
		from:  from,
		to:    to,
	}
	return inst
}

// Allocate attempts to allocate a continuous range of blocks as specified, returning
// the first blockid in the allocation, and whether the space was successfully allocated.
func (nd *Node) Allocate(blocks uint64) (uint64, bool) {
	if nd.left != nil {
		blockid, found := nd.left.Allocate(blocks)
		if found {
			return blockid, found
		}
	}
	if nd.to-nd.from >= blocks {
		// Will fit in this node.
		blockid := nd.from
		nd.from += blocks
		return blockid, true
	}
	if nd.right != nil {
		blockid, found := nd.right.Allocate(blocks)
		if found {
			return blockid, found
		}
	}

	return 0, false
}

// Deallocate frees the continuous space referenced by the specified blockid and length, returning the new node
// representing the free space.
func (nd *Node) Deallocate(blockid uint64, blocklength uint64) *Node {
	node := NewNode(blockid, blockid+blocklength)
	nd = nd.AddNode(node)
	return nd
}

// AddNode adds an existing node into the tree, merging nodes if necessary and returning the new root of the tree.
func (nd *Node) AddNode(node *Node) *Node {
	// Detect node engulfed by nd
	if nd.from <= node.from && nd.to >= node.to {
		// Add node's children if any
		if node.left != nil {
			nd.AddNode(node.left)
		}
		if node.right != nil {
			nd.AddNode(node.right)
		}
		return nd // drop node
	}
	// Detect nd engulfed by node
	if node.from <= nd.from && node.to >= nd.to {
		// add our children to new node
		if nd.left != nil {
			node.AddNode(nd.left)
		}
		if nd.right != nil {
			node.AddNode(nd.right)
		}
		// drop nd, return new node
		nd.left = nil
		nd.right = nil
		return node
	}
	// Detect adjacent to left / overlaps left
	if node.to == nd.from-1 || (node.from <= nd.from && node.to <= nd.to && node.to >= nd.from) {
		nd.from = node.from // extend nd and drop new node
		// Add node's children if any
		if node.left != nil {
			nd.AddNode(node.left)
		}
		if node.right != nil {
			nd.AddNode(node.right)
		}
		// Clear and re-add children
		left := nd.left
		right := nd.right
		nd.left = nil
		nd.right = nil
		if left != nil {
			nd.AddNode(left)
		}
		if right != nil {
			nd.AddNode(right)
		}
		// drop node
		node.left = nil
		node.right = nil
		return nd
	}
	// Detect adjacent to right / overlaps right
	if node.from == nd.to+1 || (node.from >= nd.from && node.from <= nd.to && node.to <= nd.from) {
		nd.to = node.to // extend nd
		// Add node's children if any
		if node.left != nil {
			nd.AddNode(node.left)
		}
		if node.right != nil {
			nd.AddNode(node.right)
		}
		// Clear and re-add children
		left := nd.left
		right := nd.right
		nd.left = nil
		nd.right = nil
		if left != nil {
			nd.AddNode(left)
		}
		if right != nil {
			nd.AddNode(right)
		}
		// drop node
		node.left = nil
		node.right = nil
		return nd
	}
	// else, binary insert
	if node.to < nd.from {
		if nd.left == nil {
			nd.left = node
		} else {
			nd.left = nd.left.AddNode(node)
		}
	} else {
		if nd.right == nil {
			nd.right = node
		} else {
			nd.right = nd.right.AddNode(node)
		}
	}
	return nd
}
