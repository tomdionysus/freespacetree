package freespacetree

import(
  // "fmt"
)

type Node struct {
  from uint64
  to uint64
  left *Node
  right *Node
}

func NewNode(from, to uint64) *Node {
  inst := &Node{
    left: nil,
    right: nil,
    from: from,
    to: to,
  }
  return inst
}

func (me *Node) Allocate(blocks uint64, limit uint64) (uint64, bool) {
  if me.to-me.from >= blocks {
    // Will fit in this node.
    blockid := me.from
    me.from += blocks
    return blockid, true
  }

  if me.right !=nil {
    blockid, found := me.right.Allocate(blocks, limit)
    if found { return blockid, found }
  } else {
    // Rightmost in tree.
    if limit - me.to > blocks {
      // Will fit between this node and limit.
      me.right = NewNode(me.to+1, blocks)
      return 0, true
    } else { return 0, false }
  }

  if me.left !=nil { 
    blockid, found := me.left.Allocate(blocks, limit)
    if found { return blockid, found }
  } else {
    // Leftmost in tree.
    if me.from > blocks {
      // Will fit between 0 and this node.
      me.left = NewNode(0, blocks)
      return 0, true
    } else { return 0, false }
  }

  return 0, false
}

func (me *Node) Deallocate(blockid uint64, blocklength uint64) *Node {
  node := NewNode(blockid, blockid+blocklength)
  me.AddNode(node)
  return me
}

// Add an existing node into the tree, merging if necessary, error if node overlaps
func (me *Node) AddNode(node *Node) *Node {
  if node.to < me.from {
    if me.left == nil {
      me.left = node
    } else { me.left = me.left.AddNode(node) }
  } else {
    if me.right == nil {
      me.right = node
    } else { me.left = me.right.AddNode(node) }
  }
  return me
}

func (me *Node) RemoveNode(node *Node) error {
  return nil
}
