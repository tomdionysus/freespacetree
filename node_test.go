package freespacetree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode(t *testing.T) {
	x := NewNode(5, 7)

	assert.NotNil(t, x)
	assert.Equal(t, x.from, uint64(5))
	assert.Equal(t, x.to, uint64(7))
}

func TestAllocate(t *testing.T) {
	// Striaghtforward alloc
	root := NewNode(0, 20)

	id, ok := root.Allocate(10)
	assert.Equal(t, id, uint64(0))
	assert.Equal(t, root.from, uint64(10))
	assert.Equal(t, root.to, uint64(20))
	assert.Nil(t, root.left)
	assert.Nil(t, root.right)
	assert.True(t, ok)

	id, ok = root.Allocate(10)
	assert.Equal(t, id, uint64(10))
	assert.Equal(t, root.from, uint64(20))
	assert.Equal(t, root.to, uint64(20))
	assert.Nil(t, root.left)
	assert.Nil(t, root.right)
	assert.True(t, ok)

	id, ok = root.Allocate(10)
	assert.Equal(t, id, uint64(0))
	assert.False(t, ok)
}

func TestFillAllocate(t *testing.T) {
	root := NewNode(0, 20)

	id, ok := root.Allocate(15)
	assert.Equal(t, id, uint64(0))
	assert.Equal(t, root.from, uint64(15))
	assert.Equal(t, root.to, uint64(20))
	assert.Nil(t, root.left)
	assert.Nil(t, root.right)
	assert.True(t, ok)

	id, ok = root.Allocate(5)
	assert.Equal(t, id, uint64(15))
	assert.Equal(t, root.from, uint64(20))
	assert.Equal(t, root.to, uint64(20))
	assert.Nil(t, root.left)
	assert.Nil(t, root.right)
	assert.True(t, ok)

	root = root.Deallocate(0, 5)
	assert.NotNil(t, root)
	assert.NotNil(t, root.left)
	assert.Equal(t, root.left.from, uint64(0))
	assert.Equal(t, root.left.to, uint64(5))
}

func TestAddNodeLeftAdjacentMerge(t *testing.T) {
	root := NewNode(40, 50)
	root.AddNode(NewNode(100, 110))
	node := NewNode(30, 39)
	node.AddNode(NewNode(0, 10))
	node.AddNode(NewNode(45, 50))

	root = root.AddNode(node)
	assert.Equal(t, root.from, uint64(30))
	assert.Equal(t, root.to, uint64(50))
	assert.NotNil(t, root.left)
	assert.Nil(t, root.left.left)
	assert.Nil(t, root.left.right)
	assert.NotNil(t, root.right)
}

func TestAddNodeRightAdjacentMerge(t *testing.T) {
	root := NewNode(40, 50)
	root.AddNode(NewNode(10, 20))

	node := NewNode(51, 60)
	node.AddNode(NewNode(80, 90))
	node.AddNode(NewNode(45, 46))

	root = root.AddNode(node)
	assert.Equal(t, root.from, uint64(40))
	assert.Equal(t, root.to, uint64(60))
	assert.NotNil(t, root.left)
	assert.NotNil(t, root.right)
}

func TestAddNodeEngulfs(t *testing.T) {
	root := NewNode(0, 20)
	node := NewNode(10, 15)

	root = root.AddNode(node)
	assert.Equal(t, root.from, uint64(0))
	assert.Equal(t, root.to, uint64(20))
	assert.Nil(t, root.left)
	assert.Nil(t, root.right)
}

func TestAddNodeEngulfsChildren(t *testing.T) {
	root := NewNode(0, 20)

	node := NewNode(10, 15)
	node = node.AddNode(NewNode(5, 7))
	node = node.AddNode(NewNode(18, 20))

	assert.NotNil(t, node.right)

	root = root.AddNode(node)
	assert.Equal(t, root.from, uint64(0))
	assert.Equal(t, root.to, uint64(20))
	assert.Nil(t, root.left)
	assert.Nil(t, root.right)
}

func TestAddNodeIsEngulfedNoChildren(t *testing.T) {
	root := NewNode(10, 15)
	node := NewNode(0, 25)

	root = root.AddNode(node)
	assert.Equal(t, root.from, uint64(0))
	assert.Equal(t, root.to, uint64(25))
	assert.Nil(t, root.left)
	assert.Nil(t, root.right)
}

func TestAddNodeIsEngulfedChildren(t *testing.T) {
	root := NewNode(10, 15)
	root.AddNode(NewNode(0, 1))
	root.AddNode(NewNode(22, 23))

	node := NewNode(5, 20)

	root = root.AddNode(node)
	assert.Equal(t, root.from, uint64(5))
	assert.Equal(t, root.to, uint64(20))
	assert.NotNil(t, root.left)
	assert.NotNil(t, root.right)
	assert.Equal(t, root.left.from, uint64(0))
	assert.Equal(t, root.left.to, uint64(1))
	assert.Equal(t, root.right.from, uint64(22))
	assert.Equal(t, root.right.to, uint64(23))
}

func TestAddNodeDepth(t *testing.T) {
	root := NewNode(40, 50)
	root.AddNode(NewNode(20, 30))
	root.AddNode(NewNode(0, 10))

	root.AddNode(NewNode(60, 70))
	root.AddNode(NewNode(80, 90))

	assert.NotNil(t, root.left)
	assert.NotNil(t, root.right)
	assert.NotNil(t, root.left.left)
	assert.Nil(t, root.left.right)
	assert.NotNil(t, root.right.right)
	assert.Nil(t, root.right.left)
}

func TestAddNodeDepthEngulf(t *testing.T) {
	root := NewNode(40, 50)
	root.AddNode(NewNode(0, 10))
	root.AddNode(NewNode(20, 30))

	root.AddNode(NewNode(60, 70))
	root.AddNode(NewNode(80, 90))

	root.AddNode(NewNode(0, 30))

	assert.NotNil(t, root.left)
	assert.NotNil(t, root.right)
	assert.Nil(t, root.left.left)
	assert.Nil(t, root.left.right)
	assert.NotNil(t, root.right.right)
	assert.Nil(t, root.right.left)
}

func TestAllocateLeftEnd(t *testing.T) {
	root := NewNode(40, 50)

	root.AddNode(NewNode(60, 70))
	root.AddNode(NewNode(80, 180))

	block1, ok := root.Allocate(100)
	assert.True(t, ok)
	assert.Equal(t, block1, uint64(80))

	root.Deallocate(60, 40)
}

func TestAllocateRightEnd(t *testing.T) {
	root := NewNode(40, 50)

	root.AddNode(NewNode(20, 30))
	root.AddNode(NewNode(0, 5))

	block1, ok := root.Allocate(5)
	assert.True(t, ok)
	assert.Equal(t, block1, uint64(0))

	root.Deallocate(60, 40)
}
