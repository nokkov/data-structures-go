package linked_list

type Singly[T any] struct {
	length int
	Head   Node[T]
}
