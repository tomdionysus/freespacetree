package freespacetree

type Node struct {
	from  uint64
	to    uint64
	left  *Node
	right *Node
}

func NewNode(from, to uint64) *Node {
	inst := &Node{
		left:  nil,
		right: nil,
		from:  from,
		to:    to,
	}
	return inst
}

func (me *Node) Allocate(blocks uint64) (uint64, bool) {
	if me.left != nil {
		blockid, found := me.left.Allocate(blocks)
		if found {
			return blockid, found
		}
	}
	if me.to-me.from >= blocks {
		// Will fit in this node.
		blockid := me.from
		me.from += blocks
		return blockid, true
	}
	if me.right != nil {
		blockid, found := me.right.Allocate(blocks)
		if found {
			return blockid, found
		}
	}

	return 0, false
}

func (me *Node) Deallocate(blockid uint64, blocklength uint64) *Node {
	node := NewNode(blockid, blockid+blocklength)
	me.AddNode(node)
	return me
}

// Add an existing node into the tree, merging if necessary
// Returns the new root of the tree
func (me *Node) AddNode(node *Node) *Node {
	// Detect node engulfed by me
	if me.from <= node.from && me.to >= node.to {
		// Add node's children if any
		if node.left != nil {
			me.AddNode(node.left)
		}
		if node.right != nil {
			me.AddNode(node.right)
		}
		return me // drop node
	}
	// Detect me engulfed by node
	if node.from <= me.from && node.to >= me.to {
		// add our children to new node
		if me.left != nil {
			node.AddNode(me.left)
		}
		if me.right != nil {
			node.AddNode(me.right)
		}
		// drop me, return new node
		me.left = nil
		me.right = nil
		return node
	}
	// Detect adjacent to left / overlaps left
	if node.to == me.from-1 || (node.from <= me.from && node.to <= me.to && node.to >= me.from) {
		me.from = node.from // extend me and drop new node
		// Add node's children if any
		if node.left != nil {
			me.AddNode(node.left)
		}
		if node.right != nil {
			me.AddNode(node.right)
		}
		// Clear and re-add children
		left := me.left
		right := me.right
		me.left = nil
		me.right = nil
		if left != nil {
			me.AddNode(left)
		}
		if right != nil {
			me.AddNode(right)
		}
		// drop node
		node.left = nil
		node.right = nil
		return me
	}
	// Detect adjacent to right / overlaps right
	if node.from == me.to+1 || (node.from >= me.from && node.from <= me.to && node.to <= me.from) {
		me.to = node.to // extend me
		// Add node's children if any
		if node.left != nil {
			me.AddNode(node.left)
		}
		if node.right != nil {
			me.AddNode(node.right)
		}
		// Clear and re-add children
		left := me.left
		right := me.right
		me.left = nil
		me.right = nil
		if left != nil {
			me.AddNode(left)
		}
		if right != nil {
			me.AddNode(right)
		}
		// drop node
		node.left = nil
		node.right = nil
		return me
	}
	// else, binary insert
	if node.to < me.from {
		if me.left == nil {
			me.left = node
		} else {
			me.left = me.left.AddNode(node)
		}
	} else {
		if me.right == nil {
			me.right = node
		} else {
			me.right = me.right.AddNode(node)
		}
	}
	return me
}
