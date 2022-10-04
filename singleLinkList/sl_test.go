package linklist

import "testing"

func IntCmp(a, b interface{}) int {
	return a.(int) - b.(int)
}
func TestNewList(t *testing.T) {
	d := NewList(IntCmp)
	if d.head != nil || d.cmp == nil {
		t.Errorf("at least one of [d.first d.last d.cmp] is nil")
	}
}

func TestAddAtTail(t *testing.T) {
	d := NewList(IntCmp)
	d.AddAtTail(1)
	d.AddAtTail(2)
	d.AddAtTail(3)
	want := 1

	if d.head.Val.(int) != want {
		t.Errorf("d.last val got %d, want:%d", d.head.Val.(int), want)
	}

}

func TestAddAtHead(t *testing.T) {
	d := NewList(IntCmp)

	d.AddAtHead(1)
	d.AddAtHead(2)
	d.AddAtHead(3)
	want := 3

	if d.head.Val.(int) != want {
		t.Errorf("d.last val got %d, want:%d", d.head.Val.(int), want)
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
