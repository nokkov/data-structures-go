package main

func main() {
	var list *LinkedList = &LinkedList{}
	list.Insert(20)
	list.ConvertListToCircular()
	list.ShowLinkedList()
}
