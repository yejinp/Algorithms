package stack

import "testing"

func TestStackEmpty(t *testing.T) {
	s := StackInit(1)

	if !s.StackEmpty() {
		t.Errorf("s.StackEmpty() got %v, want: %v", false, true)
	}

	s = StackInit(0)
	if !s.StackEmpty() {
		t.Errorf("s.StackEmpty() got %v, want: %v", true, false)
	}
}

func TestStackFull(t *testing.T) {

}
