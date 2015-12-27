package freespacetree

type Tree struct {
  root *Node
  capacity uint64
}

func New(capacity uint64) *Tree {
  inst := &Tree{
    root: NewNode(0, capacity),
    capacity: capacity,
  }
  return inst
}

func (me *Tree) Allocate(blocks uint64) (uint64, bool) {
  return me.root.Allocate(blocks)
}

func (me *Tree) Deallocate(blockid uint64, blocklength uint64) {
  me.root = me.root.Deallocate(blockid, blocklength)
}