package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	info interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(v interface{}) {
	node := &Node{info: v, next: nil}

	if l.head == nil {
		l.head = node
		return
	}

	cNode := l.head
	for cNode.next != nil {
		cNode = cNode.next
	}

	cNode.next = node
}

func (l *LinkedList) ShowLinkedList() {
	if l.head != nil {
		node := l.head

		for node != nil {
			fmt.Printf("-> %d ", node.info)
			node = node.next
		}
	}
}

func InitLinkedListExample() *LinkedList {
	head := &Node{5, nil}
	linkedList := &LinkedList{head: head}

	for i := 0; i < 10; i++ {
		linkedList.Insert(rand.Intn(11))
	}

	return linkedList
}
