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
