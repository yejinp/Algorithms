package binarytree

import (
	"fmt"
	"testing"
)

func IntTreeCmp(a, b interface{}) int {
	return int(a.(IntTree) - b.(IntTree))
}

type IntTree int

func (i *IntTree) Walk(n *Node) {
	fmt.Printf(" %d", n.val.(int))
}

func TestBinaryInsert(t *testing.T) {
	b := NewBinaryTree(IntTreeCmp)
	b.Insert(IntTree(2))
	b.Insert(IntTree(1))
	b.Insert(IntTree(3))
	if b.Root == nil || b.Root.val != IntTree(2) {
		t.Errorf("b.root.val = %d, want:%d", b.Root.val.(IntTree), 2)
	}
	if b.Root.left == nil || b.Root.left.val != IntTree(1) {
		t.Errorf("b.root.left.val = %d, want:%d", b.Root.left.val.(IntTree), 1)
	}
	if b.Root.right == nil || b.Root.right.val != IntTree(3) {
		t.Errorf("b.root.right.val = %d, want:%d", b.Root.right.val.(IntTree), 3)
	}
}

func TestBinaryDelete(t *testing.T) {
	b := NewBinaryTree(IntTreeCmp)
	b.Insert(IntTree(2))
	b.Insert(IntTree(1))
	b.Insert(IntTree(3))

	d := b.Minimum(b.Root)
	b.Delete(d)
	if b.Root.left != nil {
		t.Errorf("b.Root.left != nil")
	}
	d = b.Maximum(b.Root)
	if d.val != IntTree(3) {
		t.Errorf("Maximum of b != %d", 3)
	}
	b.Delete(d)
	if b.Root.right != nil {
		t.Errorf("b.Root.right != nil")
	}
}

func TestBinaryWalk(t *testing.T) {
	b := NewBinaryTree(IntTreeCmp)
	b.Insert(IntTree(2))
	b.Insert(IntTree(1))
	b.Insert(IntTree(3))
	b.Insert(IntTree(4))
	b.Insert(IntTree(6))

	msg, want := "", "1\n2\n3\n4\n6\n"
	InOrderWalk(b.Root, func(n *Node) error {
		msg = msg + fmt.Sprintf("%d\n", n.val.(IntTree))
		return nil
	})

	if msg != want {
		t.Errorf("walker get:[%s], want:[%s]", msg, want)
	}

}
