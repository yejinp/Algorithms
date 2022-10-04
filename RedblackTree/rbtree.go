package rbtree

type cmpFunc func(interface{}, interface{}) int

type Rbtree struct {
	Root, nill *Node
	cmp        cmpFunc
}

type Node struct {
	left, right, parent *Node
	Val                 interface{}
	color               Rb_color
}

type Rb_color uint8

const (
	RB_RED Rb_color = iota
	RB_BLACK
)

var rbtree_nil_node = Node{color: RB_BLACK}

func (t *Rbtree) left_rotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != t.nill {
		y.left.parent = x
	}
	y.parent = x.parent

	if t.nill == x.parent {
		t.Root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (t *Rbtree) right_rotate(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != t.nill {
		x.left.parent = x
	}
	y.parent = x.parent

	if x == t.Root {
		t.Root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.right = x
	x.parent = y
}

func (t *Rbtree) Insert(val interface{}) {
	z := &Node{Val: val}
	x, y := t.Root, t.nill

	for x != t.nill {
		y = x
		if t.cmp(z.Val, x.Val) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == t.nill {
		t.Root = z
	} else if t.cmp(z.Val, y.Val) < 0 {
		y.left = z
	} else {
		y.right = z
	}
	z.left = t.nill
	z.right = t.nill
	z.color = RB_RED
	t.insertFixup(z)
}

func (b *Rbtree) Minimum(x *Node) *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}
func (t *Rbtree) insertFixup(z *Node) {
	for z.parent.color == RB_RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == RB_RED {
				z.parent.color = RB_BLACK
				y.color = RB_BLACK
				z.parent.parent.color = RB_RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.left_rotate(z)
				}
				z.parent.color = RB_BLACK
				z.parent.parent.color = RB_RED
				t.right_rotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == RB_RED {
				z.parent.color = RB_BLACK
				y.color = RB_BLACK
				z.parent.parent.color = RB_RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.right_rotate(z.parent.parent)
				}
				z.parent.color = RB_BLACK
				z.parent.parent.color = RB_RED
				t.left_rotate(z.parent.parent)
			}
		}
	}
	t.Root.color = RB_BLACK
}
func (t *Rbtree) Delete(z *Node) {
	x, y := z, z
	y_original_color := y.color

	if z.left == t.nill {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t.nill {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.Minimum(z.right)
		y_original_color = y.color
		x = y.right
		if y != x.right {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		} else {
			x.parent = y
		}

		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if y_original_color == RB_BLACK {
		t.deleteFixup(x)
	}
}

func (t *Rbtree) deleteFixup(x *Node) {
	for x != t.Root && x.color == RB_BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RB_RED {
				w.color = RB_BLACK
				x.parent.color = RB_RED
				t.left_rotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == RB_BLACK && w.right.color == RB_BLACK {
				w.color = RB_RED
				x = x.parent
			} else {
				if w.right.color == RB_BLACK {
					w.left.color = RB_BLACK
					w.color = RB_RED
					t.right_rotate(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = RB_BLACK
				w.right.color = RB_BLACK
				t.left_rotate(x.parent)
				x = t.Root
			}
		} else {
			w := x.parent.left
			if w.color == RB_RED {
				w.color = RB_BLACK
				x.parent.color = RB_RED
				t.right_rotate(x.parent)
				w = x.parent.left
			}
			if w.right.color == RB_BLACK && w.left.color == RB_BLACK {
				w.color = RB_RED
				x = x.parent
			} else {
				if w.left.color == RB_BLACK {
					w.right.color = RB_BLACK
					w.color = RB_RED
					t.left_rotate(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = RB_BLACK
				w.left.color = RB_BLACK
				t.right_rotate(x.parent)
				x = t.Root
			}
		}
	}
	x.color = RB_BLACK
}

func (t *Rbtree) transplant(u, v *Node) {
	if u.parent == t.nill {
		t.Root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func NewRbtree(cmp cmpFunc) *Rbtree {
	root := Rbtree{Root: &rbtree_nil_node, nill: &rbtree_nil_node, cmp: cmp}
	return &root
}

func (t *Rbtree) Search(x *Node, key interface{}) *Node {
	for x != t.nill {
		//	fmt.Println(x, x.Val, key)
		diff := t.cmp(key, x.Val)
		if diff == 0 {
			break
		}
		if diff < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (t *Rbtree) Maximum(x *Node) *Node {
	for x != nil && x.right != nil {
		x = x.right
	}
	return x
}

func (t *Rbtree) IsEmpty() bool {
	return t.Root == t.nill
}

func (t *Rbtree) InOrderWalk(n *Node, walker func(*Node) error) {
	if walker == nil {
		return
	}
	if n != t.nill {
		t.InOrderWalk(n.left, walker)
		walker(n)
		t.InOrderWalk(n.right, walker)
	}
}
