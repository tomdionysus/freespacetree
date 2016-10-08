# freespacetree

[![Build Status](https://travis-ci.org/tomdionysus/freespacetree.svg)](https://travis-ci.org/tomdionysus/freespacetree)
[![Coverage Status](https://coveralls.io/repos/tomdionysus/freespacetree/badge.svg?branch=master&service=github)](https://coveralls.io/github/tomdionysus/freespacetree?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomdionysus/freespacetree)](https://goreportcard.com/report/github.com/tomdionysus/freespacetree)
[![GoDoc](https://godoc.org/github.com/tomdionysus/freespacetree?status.svg)](https://godoc.org/github.com/tomdionysus/freespacetree)

freespacetree is an abstract data type designed to provide the following:

* Efficiently store a representation of free space on a storage device
* Find the first available block of free space larger than *n* in polynomial time
* Allow efficient modifications (allocation/deallocation) of free space

## Overview

A storage device is assumed to be a continious range of values of type **int64** (blocks, bytes, runes, etc), with a defined start and end value. Areas of the device can be allocated, removing them from the free space pool, and previously allocated areas can be fully or partially deallocated, that is, returned to the free space pool. This is a common problem in filesystems, databases and the like, where modern disk sizes in the Gb and Tb range now mean that approaches such as a [Free Space Bitmap][1] are largely unworkable.  

As such, freespacetree attempts to solve the problem by storing representations of continuous areas of free space in a binary tree structure. 

**NOTE**: The nodes represent free space, not allocated space. As such, the node coverage is a 'negative' of usage.

## Structure

The functionality of `freespacetree` in contained within two types, `Node` and `FreeSpaceTree`.

For clarity, it is assumed that the device contains blocks, each atomic unit of free space is assumed to have a unique address/offset/identifier of type `uint64` which will be referred to as the **blockid**. There is no reason, however, that freespacetree couldn't be used for memory allocation (blockid = memory address) or for space partition (blockid = physical measurement), etc.

### Node

A `Node` represents a continuous area of free space, and has the following properties:

| Property           | Type       | Description                                                                        |
|:-------------------|:-----------|:-----------------------------------------------------------------------------------|
| from               | `uint64`   | The lowest blockid in the range                                                    |
| to                 | `uint64`   | The highest blockid in the range                                                   |
| left               | `*Node`    | The root `Node` of a tree containing all `Node`s where `to` < this `Node`'s `from` |
| right              | `*Node`    | The root `Node` of a tree containing all `Node`s where `from` > this `Node`'s `to` |

The rules for nodes within a tree are as follows:

* If the tree has no allocations, a single `Node` exists at the root, spanning the available capacity of the tree.
* All nodes to the left of any node must represent space 'lower' than the bounds of that node.
* All nodes to the right of any node must represent space 'higher' than the bounds of that node.
* A node may not overlap, engulf, or be engulfed by any other node, a merge should occur instead.
* No node may be exactly adjacent to any other node (two nodes with no allocation between them) - a merge should occur instead.

### FreeSpaceTree

The `FreeSpaceTree` type is a container for the tree, with the following properties:

| Property           | Type       | Description                                      |
|:-------------------|:-----------|:-------------------------------------------------|
| capacity           | `uint64`   | The largest blockid that the device can address  |
| free               | `uint64`   | The total free blocks in the tree                |
| root               | `*Node`    | The root `Node` of the tree                      |

## Operations

### Allocation

Allocation is the process where *n* continuous free blocks are requested from the tree, which attempts to find a suitable area of free space and allocate it. The response is one of the following:

* Returns the blockid of the first block in the area, having allocated the free space.
* No continuous area of free space is available of the size requested.

### Deallocation

Allocation is the process where *n* continuous free blocks are freed from the tree, which marks those blocks as free and returns accordingly.
As allocations are 'valueless', i.e. there is no specific information associated with the allocation other than that area of blocks is now unavailable, the tree will merge adjacent blocks of free space created during deallocation in order to maintain storage and operational efficiency.

## License

freespacetree is licensed under the Open Source MIT license. Please see the [License File](LICENSE.txt) for more details.

## Code Of Conduct

The freespacetree project supports and enforces [The Contributor Covenant](http://contributor-covenant.org/). Please read [the code of conduct](CODE_OF_CONDUCT.md) before contributing.

## Further Reading

[1]: (https://en.wikipedia.org/wiki/Free_space_bitmap)