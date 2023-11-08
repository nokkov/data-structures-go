package linked_list

type Singly[T any] struct {
	length int
	Head   *Node[T]
}

func NewSingly[T any]() *Singly[T] {
	return &Singly[T]{}
}

func (sl *Singly[T]) AddAtBeg(val T) {
	node := NewNode(val)
	node.next = sl.Head
	sl.Head = node
	sl.length++
}

func (sl *Singly[T]) AddAtEnd(val T) {
	node := NewNode(val)
	head := sl.Head

	if head != nil {

		for head.next != nil {
			head = head.next
		}

		head.next = node
	} else {
		sl.Head = node
	}

	sl.length++
}

// DeleteAtBeg delete head node
// return its value and false if list was empty
func (sl *Singly[T]) DeleteAtBeg() (T, bool) {
	if sl.Head == nil {
		var r T
		return r, false
	}

	cur := sl.Head
	sl.Head = cur.next
	sl.length--

	return cur.info, true
}
