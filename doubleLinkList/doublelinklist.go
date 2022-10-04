package doublelinklist

type cmpFunc func(interface{}, interface{}) int

type DoubleList struct {
	first, last *Node
	cmp         cmpFunc //compare function used to compare the val in Node
}

type Node struct {
	next, prev *Node
	Val        interface{}
}

func NewList(cmp cmpFunc) *DoubleList {
	list := DoubleList{}
	list.cmp = cmp
	return &list
}

func (d *DoubleList) AddAtTail(val interface{}) {
	last := d.last
	node := &Node{Val: val}

	if d.first == nil {
		d.first = node
		d.last = node
	} else {
		last.next = node
		node.prev = last
		d.last = node
	}
}

func (d *DoubleList) InOrderWalk(node *Node, walker func(*Node) error) {
	for node != d.last {
		walker(node)
		node = node.next
	}
}

func (d *DoubleList) AddAtHead(val interface{}) {
	first := d.first
	node := &Node{Val: val}

	if d.first == nil {
		d.first = node
		d.last = node
	} else {
		first.prev = node
		node.next = first
		d.first = node
	}
}

func (d *DoubleList) Search(val interface{}) *Node {
	node := d.first

	for node != nil {
		if d.cmp(node.Val, val) == 0 {
			break
		}
		node = node.next
	}
	return node
}
