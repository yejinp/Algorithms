// Copyright 2022 yejinping. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binarytree provides primitives for binary Tree.
package binarytree

type cmpFunc func(interface{}, interface{}) int

type BinaryTree struct {
	Root *Node
	cmp  cmpFunc //compare function used to compare the val in Node
}

type Node struct {
	left, right, parent *Node
	val                 interface{} // empty interface, which can store any type
}

func (b *BinaryTree) Insert(val interface{}) {
	z := &Node{val: val}
	x, y := b.Root, b.Root

	for x != nil {
		y = x
		if b.cmp != nil && b.cmp(val, x.val) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y

	if y == nil {
		b.Root = z
	} else if b.cmp != nil && b.cmp(z.val, y.val) < 0 {
		y.left = z
	} else {
		y.right = z
	}
}

func (b *BinaryTree) Search(x *Node, key interface{}) *Node {
	if x == nil {
		return x
	}
	diff := b.cmp(key, x.val)
	if diff == 0 {
		return x
	}
	if diff < 0 {
		return b.Search(x.left, key)
	}

	return b.Search(x.right, key)
}

func (b *BinaryTree) Iterative_Search(x *Node, key interface{}) *Node {
	for x != nil {
		diff := b.cmp(key, x.val)
		if diff == 0 {
			break
		}
		if diff < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (b *BinaryTree) Minimum(x *Node) *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}

func (b *BinaryTree) Maximum(x *Node) *Node {
	for x != nil && x.right != nil {
		x = x.right
	}
	return x
}

func InOrderWalk(n *Node, walker func(*Node) error) {
	if walker == nil {
		return
	}
	if n != nil {
		InOrderWalk(n.left, walker)
		walker(n)
		InOrderWalk(n.right, walker)
	}
}

func (b *BinaryTree) Successor(x *Node) *Node {
	if x.right != nil {
		return b.Minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

func (b *BinaryTree) Predecessor(x *Node) *Node {
	if x.left != nil {
		return b.Maximum(x.left)
	}
	y := x.parent
	for y != nil && x == y.left {
		x = y
		y = y.left
	}
	return y
}

func (b *BinaryTree) Delete(z *Node) {
	if z.left == nil {
		b.transplant(z, z.right)
	} else if z.right == nil {
		b.transplant(z, z.left)
	} else {
		y := b.Minimum(z.right)
		if y != z.right {
			b.transplant(y, y.right)
			y.right = z.right
			y.right.parent = z
		}
		b.transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}
}

func (b *BinaryTree) transplant(u, v *Node) {
	if u.parent == nil {
		b.Root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func NewBinaryTree(cmp cmpFunc) *BinaryTree {
	if cmp == nil {
		return nil
	}

	tree := BinaryTree{cmp: cmp}
	return &tree
}
