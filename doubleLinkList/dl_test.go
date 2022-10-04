package doublelinklist

import "testing"

func IntCmp(a, b interface{}) int {
	return a.(int) - b.(int)
}
func TestNewList(t *testing.T) {
	d := NewList(IntCmp)
	if d.first != nil || d.last != nil || d.cmp == nil {
		t.Errorf("at least one of [d.first d.last d.cmp] is nil")
	}
}

func TestAddAtTail(t *testing.T) {
	d := NewList(IntCmp)
	d.AddAtTail(1)
	d.AddAtTail(2)
	d.AddAtTail(3)
	if d.last.Val.(int) != 3 {
		t.Errorf("d.last val got %d, want:%d", d.last.Val.(int), 3)
	}
	if d.first.Val.(int) != 1 {
		t.Errorf("d.first val got %d, want:%d", d.last.Val.(int), 1)
	}
}

func TestAddAtHead(t *testing.T) {
	d := NewList(IntCmp)

	d.AddAtHead(1)
	d.AddAtHead(2)
	d.AddAtHead(3)

	if d.last.Val.(int) != 1 {
		t.Errorf("d.last val got %d, want:%d", d.last.Val.(int), 1)
	}
	if d.first.Val.(int) != 3 {
		t.Errorf("d.first val got %d, want:%d", d.last.Val.(int), 3)
	}
}
func TestSearch(t *testing.T) {
	d := NewList(IntCmp)
	d.AddAtTail(1)
	d.AddAtTail(2)
	d.AddAtTail(3)

	want := 3
	got := d.Search(3)
	if got == nil {
		t.Errorf("d.Search(3) got nil")
	}
	if got.Val.(int) != want {
		t.Errorf("got val  %d, want:%d", got.Val.(int), want)
	}

	got = d.Search(30)
	if got != nil {
		t.Errorf("d.Search(30) got %v, want nil", got)
	}
}

func TestTraverse(t *testing.T) {
	d := NewList(IntCmp)
	d.AddAtTail(1)
	d.AddAtTail(2)
	d.AddAtTail(3)

	first := d.Search(1)
	for first != nil {
		second := first.next
		if second != nil {
			if first.next != second || second.prev != first {
				t.Errorf("first.next != second || second.last != first")
			}
		}
		first = first.next
	}
}
