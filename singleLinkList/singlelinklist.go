package linklist

type cmpFunc func(interface{}, interface{}) int

type SingleList struct {
	head *Node
	cmp  cmpFunc //compare function used to compare the val in Node
}

type Node struct {
	next *Node
	Val  interface{}
}

func NewList(cmp cmpFunc) *SingleList {
	list := SingleList{cmp: cmp}
	return &list
}

func (l *SingleList) AddAtTail(val interface{}) {
	head, tail := l.head, l.head
	node := Node{Val: val}
	if l.head == nil {
		l.head = &node
		return
	}
	for head != nil {
		tail = head
		head = head.next
	}
	tail.next = &node
}

func (l *SingleList) AddAtHead(val interface{}) {
	node := Node{next: l.head, Val: val}
	l.head = &node
}

func (l *SingleList) Search(val interface{}) *Node {
	node := l.head

	for node != nil {
		if l.cmp(node.Val, val) == 0 {
			break
		}
		node = node.next
	}
	return node
}
