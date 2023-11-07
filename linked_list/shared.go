package linked_list

// Node represent the linked list node
// This node is shared across different implementations
type Node[T any] struct {
	info T
	next *Node[T]
	prev *Node[T]
}

// NewNode Create new node (wow)
func NewNode[T any](val T) *Node[T] {
	return &Node[T]{val, nil, nil}
}
