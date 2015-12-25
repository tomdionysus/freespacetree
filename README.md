# freespacetree

**CAVEAT**: freespacetree is currently in early alpha. Please do not use freespacetree at this time.

[![Build Status](https://travis-ci.org/tomdionysus/freespacetree.svg)](https://travis-ci.org/tomdionysus/freespacetree)
[![Coverage Status](https://coveralls.io/repos/tomdionysus/freespacetree/badge.svg?branch=master&service=github)](https://coveralls.io/github/tomdionysus/freespacetree?branch=master)
[![GoDoc](https://godoc.org/github.com/tomdionysus/freespacetree?status.svg)](https://godoc.org/github.com/tomdionysus/freespacetree)

freespacetree is an abstract data type designed to provide the following:

* Efficiently store a representation of free space on a storage device
* Find the first available block of free space larger than *n* in polynomial time
* Allow efficient modifications (allocation/deallocation) of free space

## Overview

A storage device is assumed to be a continious range of values of type **int64** (blocks, bytes, runes, etc), with a defined start and end value. Areas of the device can be allocated, removing them from the free space pool, and previously allocated areas can be fully or partially deallocated, that is, returned to the free space pool. This is a common problem in filesystems, databases and the like, where modern disk sizes in the Gb and Tb range now mean that approaches such as a [Free Space Bitmap][1] are largely unworkable.  

As such, freespacetree attempts to solve the problem by storing representations of continuous areas of free space in a binary tree structure.

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

### FreeSpaceTree

## Operations

### Allocation

### Deallocation

## Further Reading

[1]: (https://en.wikipedia.org/wiki/Free_space_bitmap#Advanced_techniques)