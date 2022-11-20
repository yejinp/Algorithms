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
	s := StackInit(0)
	if !s.StackFull() {
		t.Errorf("s.StackFull() got %v, want: %v", true, false)
	}
	s = StackInit(1)
	if s.StackFull() {
		t.Errorf("s.StackFull() got %v, want: %v", true, false)
	}

}

func TestStackPush(t *testing.T) {
	s := StackInit(0)
	if err := s.Push(1); err == nil {
		t.Errorf("s.Push(1) got %v", err)
	}

	s = StackInit(1)
	if err := s.Push(1); err != nil {
		t.Errorf("s.Push(1) got %v", err)
	}
}

func TestStackSize(t *testing.T) {
	s := StackInit(6)
	if size := s.Size(); size != 6 {
		t.Errorf("s.Size() got %v", size)
	}

	s = StackInit(1)
	if size := s.Size(); size != 1 {
		t.Errorf("s.Size() got %v", size)
	}
}

func TestStackPop(t *testing.T) {
	s := StackInit(8)

	if err, _ := s.Pop(); err == nil {
		t.Errorf("s.Pop() got %v", err)
	}

	if err := s.Push(1); err != nil {
		t.Errorf("s.Push(1) got %v", err)
	}

	if err, val := s.Pop(); err != nil || val != 1 {
		t.Errorf("s.Pop() got %v,val:%v", err, val)
	}
}
