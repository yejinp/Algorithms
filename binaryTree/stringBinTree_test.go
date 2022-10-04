package binarytree

import (
	"fmt"
	"strings"
	"testing"
)

func stringCmp(a, b interface{}) int {
	x := a.(string)
	y := b.(string)

	return strings.Compare(x, y)
}

func TestStringBinaryInsert(t *testing.T) {
	b := NewBinaryTree(stringCmp)
	b.Insert("2")
	b.Insert("1")
	b.Insert("3")
	if b.Root == nil || b.Root.val.(string) != "2" {
		t.Errorf("b.root.val = %s, want:%s", b.Root.val.(string), "2")
	}
	if b.Root.left == nil || b.Root.left.val.(string) != "1" {
		t.Errorf("b.root.left.val = %s, want:%s", b.Root.left.val.(string), "1")
	}
	if b.Root.right == nil || b.Root.right.val.(string) != "3" {
		t.Errorf("b.root.right.val = %s, want:%s", b.Root.right.val.(string), "3")
	}
}

func TestStringBinaryDelete(t *testing.T) {
	b := NewBinaryTree(stringCmp)
	b.Insert("2")
	b.Insert("1")
	b.Insert("3")

	d := b.Minimum(b.Root)
	b.Delete(d)
	if b.Root.left != nil {
		t.Errorf("b.Root.left != nil")
	}
	d = b.Maximum(b.Root)
	if d.val.(string) != "3" {
		t.Errorf("Maximum of b want %s, got:%s", "3", d.val.(string))
	}
	b.Delete(d)
	if b.Root.right != nil {
		t.Errorf("b.Root.right != nil")
	}
}

func TestStringBinaryWalk(t *testing.T) {
	b := NewBinaryTree(stringCmp)
	b.Insert("9")
	b.Insert("1")
	b.Insert("3")
	b.Insert("4")
	b.Insert("6")

	msg, want := "", "1\n3\n4\n6\n9\n"
	InOrderWalk(b.Root, func(n *Node) error {
		msg = msg + fmt.Sprintf("%s\n", n.val.(string))
		return nil
	})

	if msg != want {
		t.Errorf("walker get:[%s], want:[%s]", msg, want)
	}

}
