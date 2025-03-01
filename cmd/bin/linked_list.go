package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// O(1) we assign a new head and point next to previous head
func (l *LinkedList) Prepend(data int) {
	node := Node{data: data}

	tmpHead := l.head
	l.head = &node
	l.head.next = tmpHead
}

// O(n) since we need to traverse the list to append a node
func (l *LinkedList) Append(data int) {
	node := Node{data: data}

	if l.head == nil {
		l.head = &node

		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = &node
}

// O(n) since we need to traverse the list find the value
func (l *LinkedList) Delete(data int) {
	if l.head == nil {
		return
	}

	if l.head.data == data {
		l.head = l.head.next
	}

	current := l.head

	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}

	// prev.next = current.next.next
}

func (l *LinkedList) PrintList() {
	current := l.head

	for current.next != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	fmt.Println(current.data, "->", nil)
}
