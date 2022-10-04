package rbtree

import "testing"

func IntCmp(a, b interface{}) int {
	return a.(int) - b.(int)
}

func TestIntRBTreeInsert(t *testing.T) {
	b := NewRbtree(IntCmp)

	b.Insert(2)
	if b.Root.color != RB_BLACK {
		t.Errorf("b.root.color != RB_BLACK, want BLACK after insert 2")
	}

	b.Insert(1)
	if b.Root.color != RB_BLACK {
		t.Errorf("b.root.color != RB_BLACK, want BLACK after insert 2 and 1")
	}

	if b.Root.left.color != RB_RED {
		t.Errorf("b.Root.left.color != RB_RED, want BLACK after insert 2 and 1")
	}

	b.Insert(3)
	if b.Root.color != RB_BLACK {
		t.Errorf("b.root.color != RB_BLACK, want BLACK after insert 2 1 and 3")
	}

	if b.Root.right.color != RB_RED {
		t.Errorf("b.Root.right.color != RB_RED, want BLACK after insert 2 1 and 3")
	}

	if b.Root == b.nill || b.Root.Val != 2 {
		t.Errorf("b.root.val = %d, want:%d", b.Root.Val.(int), 2)
	}

	if b.Root.left == b.nill || b.Root.left.Val != 1 {
		t.Errorf("b.root.left.val = %d, want:%d", b.Root.left.Val.(int), 1)
	}

	if b.Root.right == b.nill || b.Root.right.Val != 3 {
		t.Errorf("b.root.right.val = %d, want:%d", b.Root.right.Val.(int), 3)
	}

	if b.Root.color != RB_BLACK {
		t.Errorf(" b.root.color != RB_BLACK, want BLACK")
	}
}

func TestIntRBTreeDelete(t *testing.T) {
	b := NewRbtree(IntCmp)

	b.Insert(2)
	b.Insert(1)
	b.Insert(3)

	r := b.Root
	delNode := b.Search(r, 3)
	b.Delete(delNode)

	if b.IsEmpty() {
		t.Errorf("After delete only data 3, rbtree is empty")
	}
	delNode = b.Search(r, 2)
	b.Delete(delNode)

	delNode = b.Search(r, 1)
	b.Delete(delNode)
	if !b.IsEmpty() {
		t.Errorf("After delete all data, rbtree is no empty")
	}
}

func TestIntRBTreeSearch(t *testing.T) {
	b := NewRbtree(IntCmp)
	b.Insert(9)
	b.Insert(1)
	b.Insert(3)
	b.Insert(4)
	b.Insert(6)

	want := 6
	node := b.Search(b.Root, want)
	if node == b.nill {
		t.Errorf("Search %d get nil", want)
	} else if node.Val.(int) != want {
		t.Errorf("Iterative_Search %d, got:%d", want, node.Val.(int))
	}

	want = 19
	node = b.Search(b.Root, want)
	if node != b.nill {
		t.Errorf("Iterative_Search %d get unnil, want nil", want)
	}
}
