package queue

import (
	"errors"
)

type Queue struct {
	tail int
	head int
	size int
	dat  []interface{}
}

func QueueInit(size int) (error, Queue) {
	if size < 3 || (size&(size-1) != 0) {
		return errors.New("overflow"), Queue{0, 0, size, make([]interface{}, size)}
	}
	q := Queue{0, 0, size, make([]interface{}, size)}
	q.size = size
	return nil, q
}

func (q *Queue) IsFull() bool {
	tail := q.tail % q.size
	head := q.head % q.size
	if (tail+1)%q.size == head {
		return true
	}
	return false
}

func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) Enqueue(x interface{}) error {
	if q.IsFull() {
		return errors.New("overflow")
	}
	q.dat[q.tail] = x
	q.tail = (q.tail + 1) % q.size
	return nil
}

func (q *Queue) Dequeue() (error, interface{}) {
	if q.IsEmpty() {
		return errors.New("underflow"), nil
	}

	x := q.dat[q.head]
	q.head = (q.head + 1) % q.size
	return nil, x
}
