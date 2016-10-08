package freespacetree

type Tree struct {
	root     *Node
	capacity uint64
}

func New(capacity uint64) *Tree {
	inst := &Tree{
		root:     NewNode(0, capacity),
		capacity: capacity,
	}
	return inst
}

func (tr *Tree) Allocate(blocks uint64) (uint64, bool) {
	return tr.root.Allocate(blocks)
}

func (tr *Tree) Deallocate(blockid uint64, blocklength uint64) {
	tr.root = tr.root.Deallocate(blockid, blocklength)
}
