package stack

import "errors"

type Stack struct {
	top  int
	size int
	dat  []interface{}
}

func StackInit(size int) Stack {
	s := Stack{0, size, make([]interface{}, size)}
	s.size = size
	return s
}

func (s *Stack) StackEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

func (s *Stack) StackFull() bool {
	return s.top == s.size
}

func (s *Stack) Push(x interface{}) error {
	if s.top == s.size {
		return errors.New("overflow")
	}

	s.dat[s.top] = x
	s.top++
	return nil
}

func (s *Stack) Pop() (error, interface{}) {
	if s.StackEmpty() {
		return errors.New("underflow"), nil
	}
	s.top = s.top - 1
	return nil, s.dat[s.top]
}

func (s *Stack) Size() int {
	return s.size
}
