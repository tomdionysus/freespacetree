package freespacetree

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
  x := NewNode()

  assert.NotNil(t, x)
}