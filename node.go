package freespacetree

import(
  "errors"
)

type Node struct {
  from uint64
  to uint64
  left *Node
  right *Node
}

func NewNode() *Node {
  inst := &Node{}
  return inst
}

// Add an existing node into the tree, merging if necessary, error if node overlaps
func (me *Node) AddNode(node *Node) (*Node, error) {
  // Detect node merge
  if me.left == nil && node.from<me.from && node.to == me.from-1 {
    // Node is right before us
    me.from = node.from
    return me, nil
  }
  if me.right == nil && node.from>me.to && node.from == me.to+1 {
    // Node is right after us
    me.to = node.to
    return me, nil
  }
  if me.left == nil && me.right == nil && node.from<me.from && node.to>me.to {
    // Node engulfs us
    me.from = node.from
    me.to = node.to
    return me, nil
  }
  // Node doesn't touch us, add to tree
  if node.from < me.from && node.to < me.from {
    // Node is before us
    if me.left == nil { 
      me.left = node
      return node, nil
    } else { 
      return me.left.AddNode(node)
    }
  }
  if node.from > me.to && node.to > me.to {
    // Node if after us
    if me.right == nil { 
      me.right = node
      return node, nil
    } else { 
      return me.right.AddNode(node)
    }
  } 
  // Node overlaps us
  return nil, errors.New("Node overlaps")
}

func (me *Node) RemoveNode(node *Node) error {
  return nil
}
