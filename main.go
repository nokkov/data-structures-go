package main

import list2 "dsa/linked_list"

func main() {
	var list *list2.LinkedList = &list2.LinkedList{}
	list.Insert(20)
	list.ConvertListToCircular()
	list.ShowLinkedList()
}
