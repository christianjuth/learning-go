package main

import "fmt"

type Node struct {
	next  *Node
	value string
}

type List struct {
	head *Node
}

func (l *List) add(item string) {
	newNode := &Node{
		next:  nil,
		value: item,
	}

	if l.head == nil {
		l.head = newNode
		return
	}

	crnt := l.head
	for crnt.next != nil {
		crnt = crnt.next
	}
	crnt.next = newNode
}

func (l *List) remove(value string) {
	if l.head.value == value {
		l.head = l.head.next
		return
	}

	crnt := l.head
	for crnt.next != nil {
		if crnt.next.value == value {
			crnt.next = crnt.next.next
			return
		}
		crnt = crnt.next
	}
}

func (l *List) print() {
	crnt := l.head
	for crnt.next != nil {
		fmt.Print(crnt.value + ", ")
		crnt = crnt.next
	}
	fmt.Print(crnt.value)
	fmt.Println()
}

func main() {
	list := List{}

	list.add("one")
	list.add("two")
	list.add("three")
	list.add("four")

	list.remove("one")

	list.print()
}
