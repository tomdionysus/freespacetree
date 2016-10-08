package freespacetree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeNew(t *testing.T) {
	x := New(40)

	assert.NotNil(t, x)
}

func TestTreeAllocate(t *testing.T) {
	x := New(40)

	bid, ok := x.Allocate(5)
	assert.True(t, ok)
	assert.Equal(t, uint64(0), bid)

	bid, ok = x.Allocate(5)
	assert.True(t, ok)
	assert.Equal(t, uint64(5), bid)

	bid, ok = x.Allocate(10)
	assert.True(t, ok)
	assert.Equal(t, uint64(10), bid)

	bid, ok = x.Allocate(2)
	assert.True(t, ok)
	assert.Equal(t, uint64(20), bid)

	bid, ok = x.Allocate(19)
	assert.False(t, ok)
}

func TestTreeDeallocate(t *testing.T) {
	x := New(40)

	bid, ok := x.Allocate(30)
	assert.True(t, ok)
	assert.Equal(t, uint64(0), bid)

	x.Deallocate(5, 10)

	bid, ok = x.Allocate(10)
	assert.True(t, ok)
	assert.Equal(t, uint64(5), bid)

	bid, ok = x.Allocate(2)
	assert.True(t, ok)
	assert.Equal(t, uint64(30), bid)

	bid, ok = x.Allocate(19)
	assert.False(t, ok)
}
