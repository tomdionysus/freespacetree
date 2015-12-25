package freespacetree

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
  x := NewNode(5,7)

  assert.NotNil(t, x)
  assert.Equal(t, x.from, uint64(5))
  assert.Equal(t, x.to, uint64(7))
}

func TestAllocate(t *testing.T) {
  // Striaghtforward alloc
  root := NewNode(0,20)

  id, ok := root.Allocate(10,20)
  assert.Equal(t, id, uint64(0))
  assert.Equal(t, root.from, uint64(10))
  assert.Equal(t, root.to, uint64(20))
  assert.Nil(t, root.left)
  assert.Nil(t, root.right)
  assert.True(t, ok)

  id, ok = root.Allocate(10,20)
  assert.Equal(t, id, uint64(10))
  assert.Equal(t, root.from, uint64(20))
  assert.Equal(t, root.to, uint64(20))
  assert.Nil(t, root.left)
  assert.Nil(t, root.right)
  assert.True(t, ok)
  
  id, ok = root.Allocate(10,20)
  assert.Equal(t, id, uint64(0))
  assert.False(t, ok)

  // Fill Alloc
  root = NewNode(0,20)

  id, ok = root.Allocate(15,20)
  assert.Equal(t, id, uint64(0))
  assert.Equal(t, root.from, uint64(15))
  assert.Equal(t, root.to, uint64(20))
  assert.Nil(t, root.left)
  assert.Nil(t, root.right)
  assert.True(t, ok)

  id, ok = root.Allocate(5,20)
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