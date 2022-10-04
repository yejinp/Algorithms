package binarytree

import (
	"fmt"
	"testing"
)

func IntCmp(a, b interface{}) int {
	return a.(int) - b.(int)
}

func TestIntBinaryInsert(t *testing.T) {
	b := NewBinaryTree(IntCmp)
	b.Insert(2)
	b.Insert(1)
	b.Insert(3)
	if b.Root == nil || b.Root.val != 2 {
		t.Errorf("b.root.val = %d, want:%d", b.Root.val.(int), 2)
	}
	if b.Root.left == nil || b.Root.left.val != 1 {
		t.Errorf("b.root.left.val = %d, want:%d", b.Root.left.val.(int), 1)
	}
	if b.Root.right == nil || b.Root.right.val != 3 {
		t.Errorf("b.root.right.val = %d, want:%d", b.Root.right.val.(int), 3)
	}
}

func TestIntBinaryDelete(t *testing.T) {
	b := NewBinaryTree(IntCmp)
	b.Insert(2)
	b.Insert(1)
	b.Insert(3)
	b.Insert(5)

	d := b.Minimum(b.Root)
	b.Delete(d)
	if b.Root.left != nil {
		t.Errorf("b.Root.left != nil")
	}
	d = b.Maximum(b.Root)
	if d.val.(int) != 5 {
		t.Errorf("Maximum of b want %d, got:%d", 3, d.val.(int))
	}
	b.Delete(d)

	d = b.Maximum(b.Root)
	if d.val.(int) != 3 {
		t.Errorf("Maximum of b want %d, got:%d", 5, d.val.(int))
	}
}

func TestIntBinaryWalk(t *testing.T) {
	b := NewBinaryTree(IntCmp)
	b.Insert(9)
	b.Insert(1)
	b.Insert(3)
	b.Insert(4)
	b.Insert(6)

	msg, want := "", "1\n3\n4\n6\n9\n"
	InOrderWalk(b.Root, func(n *Node) error {
		msg = msg + fmt.Sprintf("%d\n", n.val.(int))
		return nil
	})

	if msg != want {
		t.Errorf("walker get:[%s], want:[%s]", msg, want)
	}
}

func TestIntBinarySearch(t *testing.T) {
	b := NewBinaryTree(IntCmp)
	b.Insert(9)
	b.Insert(1)
	b.Insert(3)
	b.Insert(4)
	b.Insert(6)

	want := 9
	node := b.Search(b.Root, want)
	if node == nil {
		t.Errorf("Search 9 get nil")
	} else if node.val.(int) != want {
		t.Errorf("Search 9 want %d, got:%d", 9, node.val.(int))
	}

	want = 19
	node = b.Search(b.Root, want)
	if node != nil {
		t.Errorf("Search %d get unnil, want nil", want)
	}
}

func TestIntBinaryIterativeSearch(t *testing.T) {
	b := NewBinaryTree(IntCmp)
	b.Insert(9)
	b.Insert(1)
	b.Insert(3)
	b.Insert(4)
	b.Insert(6)

	want := 6
	node := b.Iterative_Search(b.Root, want)
	if node == nil {
		t.Errorf("Search %d get nil", want)
	} else if node.val.(int) != want {
		t.Errorf("Iterative_Search %d, got:%d", want, node.val.(int))
	}

	want = 19
	node = b.Iterative_Search(b.Root, want)
	if node != nil {
		t.Errorf("Iterative_Search %d get unnil, want nil", want)
	}
}

func TestIntBinaryMinimum(t *testing.T) {
	b := NewBinaryTree(IntCmp)
	b.Insert(9)
	b.Insert(1)
	b.Insert(3)
	b.Insert(4)
	b.Insert(6)

	want := 1
	node := b.Minimum(b.Root)
	if node == nil {
		t.Errorf("Minimum get nil")
	} else if node.val.(int) != want {
		t.Errorf("Minimum want %d, got:%d", want, node.val.(int))
	}

}

func TestIntBinaryMaximum(t *testing.T) {
	b := NewBinaryTree(IntCmp)
	b.Insert(9)
	b.Insert(1)
	b.Insert(3)
	b.Insert(4)
	b.Insert(65)

	want := 65
	node := b.Maximum(b.Root)
	if node == nil {
		t.Errorf("Maximum get nil")
	} else if node.val.(int) != want {
		t.Errorf("Maximum want %d, got:%d", want, node.val.(int))
	}
}

func TestIntBinarySuccessor(t *testing.T) {
	b := NewBinaryTree(IntCmp)
	b.Insert(9)
	b.Insert(1)
	b.Insert(3)
	b.Insert(4)
	b.Insert(65)

	want := 1
	node := b.Iterative_Search(b.Root, want)
	if node == nil {
		t.Errorf("Iterative_Search get %d nil", want)
	} else if node.val.(int) != want {
		t.Errorf("Iterative_Search want %d, got:%d", want, node.val.(int))
	}

	want = 3
	next := b.Successor(node)
	if next == nil {
		t.Errorf("Successor %d get nil", node.val.(int))
	} else if next.val.(int) != want {
		t.Errorf("Successor of %d want %d, got:%d", node.val.(int), want, next.val.(int))
	}

	want = 65
	node = b.Iterative_Search(b.Root, want)
	if node == nil {
		t.Errorf("Iterative_Search get %d nil", want)
	} else if node.val.(int) != want {
		t.Errorf("Iterative_Search want %d, got:%d", want, node.val.(int))
	}

	next = b.Successor(node)
	if next != nil {
		t.Errorf("Successor %d get %d, want nil", node.val.(int), next.val.(int))
	}
}

func TestIntBinaryPredecessor(t *testing.T) {
	b := NewBinaryTree(IntCmp)
	b.Insert(9)
	b.Insert(1)
	b.Insert(3)
	b.Insert(4)
	b.Insert(65)

	want := 3
	node := b.Iterative_Search(b.Root, want)
	if node == nil {
		t.Errorf("Iterative_Search get %d nil", want)
	} else if node.val.(int) != want {
		t.Errorf("Iterative_Search want %d, got:%d", want, node.val.(int))
	}

	want = 1
	next := b.Predecessor(node)
	if next == nil {
		t.Errorf("Predecessor %d get nil", node.val.(int))
	} else if next.val.(int) != want {
		t.Errorf("Predecessor of %d want %d, got:%d", node.val.(int), want, next.val.(int))
	}

	want = 65
	node = b.Iterative_Search(b.Root, want)
	if node == nil {
		t.Errorf("Iterative_Search get %d nil", want)
	} else if node.val.(int) != want {
		t.Errorf("Iterative_Search want %d, got:%d", want, node.val.(int))
	}

	next = b.Predecessor(node)
	want = 9
	if next == nil {
		t.Errorf("Predecessor %d get nil", node.val.(int))
	} else if next.val.(int) != want {
		t.Errorf("Predecessor of %d want %d, got:%d", node.val.(int), want, next.val.(int))
	}
}
