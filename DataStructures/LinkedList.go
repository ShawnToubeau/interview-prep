package main

import (
	"errors"
	"fmt"
)

type LLNode struct {
	value int
	next  *LLNode
}

type LinkedList struct {
	head   *LLNode
	length int
}

// Adds to the end of the linked list.
func (list *LinkedList) addToEnd(node LLNode) {
	if list.length == 0 {
		list.head = &node
		list.length++
		return
	}

	currNode := list.head

	for currNode.next != nil {
		currNode = currNode.next
	}

	currNode.next = &node
	list.length++
}

// Adds to the front of the linked list.
func (list *LinkedList) addToStart(node LLNode) {
	headCopy := list.head
	node.next = headCopy

	list.head = &node
	list.length++
}

// Adds a node to any valid index within the list. An index of 0 inserts at the beginning,
// an index of the current list length inserts at the end.
func (list *LinkedList) addAtIndex(node LLNode, i int) error {
	// out of bounds
	if i < 0 || i > list.length {
		return errors.New("index out of bounds")
	}
	// i == 0; insert at start
	if i == 0 {
		list.addToStart(node)
		return nil
	}
	// i == list length; insert at end
	if i == list.length {
		list.addToEnd(node)
		return nil
	}

	currNode, prevNode := list.head, &LLNode{}

	for i != 0 {
		prevNode = currNode
		currNode = currNode.next
		i--
	}

	prevNode.next = &node
	node.next = currNode
	list.length++

	return nil
}

// Deletes the last node from the list.
func (list *LinkedList) deleteFromEnd() {
	if list.length == 0 {
		return
	}
	if list.length == 1 {
		list.head = nil
		list.length--
		return
	}

	currNode, prevNode := list.head, &LLNode{}

	for currNode.next != nil {
		prevNode = currNode
		currNode = currNode.next
	}

	prevNode.next = nil
	list.length--
}

// Deletes the first node from the list.
func (list *LinkedList) deleteFromStart() {
	if list.length == 0 {
		return
	}
	if list.length == 1 {
		list.head = nil
		list.length--
		return
	}

	list.head = list.head.next
	list.length--
}

// Deletes a node from any valid index within the list. An index of 0 deletes from the beginning,
// an index of the current list length deletes from the end.
func (list *LinkedList) deleteFromIndex(index int) error {
	// out of bounds
	if index < 0 || index > list.length {
		return errors.New("index out of bound")
	}
	// i == 0; delete from start
	if index == 0 {
		list.deleteFromStart()
		return nil
	}
	// i == list length; delete from end
	if index == list.length {
		list.deleteFromEnd()
		return nil
	}

	currNode, prevNode := list.head, &LLNode{}

	for index != 0 {
		prevNode = currNode
		currNode = currNode.next
		index--
	}

	prevNode.next = currNode.next
	list.length--

	return nil
}

// Prints the contents of the linked list.
func (list *LinkedList) printNodes() {
	if list.length == 0 {
		fmt.Println("List is empty")
		return
	}
	currNode := list.head

	for currNode.next != nil {
		fmt.Printf("%v ", currNode.value)
		currNode = currNode.next
	}

	fmt.Printf("%v - len %v \n", currNode.value, list.length)
}

func main() {
	list := LinkedList{}

	list.addToEnd(LLNode{value: 20})

	list.printNodes()

	list.addToStart(LLNode{value: 5})

	list.printNodes()

	list.addToEnd(LLNode{value: 29})

	list.printNodes()

	err := list.addAtIndex(LLNode{value: 42}, 2)
	if err != nil {
		fmt.Printf("Error inserting node: %v\n", err)
	}

	list.printNodes()

	err = list.addAtIndex(LLNode{value: 1}, 4)
	if err != nil {
		fmt.Printf("Error inserting node: %v\n", err)
	}

	list.printNodes()

	err = list.addAtIndex(LLNode{value: 99}, 6)
	if err != nil {
		fmt.Printf("Error inserting node: %v\n", err)
	}

	list.deleteFromEnd()

	list.printNodes()

	list.deleteFromStart()

	list.printNodes()

	err = list.deleteFromIndex(1)
	if err != nil {
		fmt.Printf("Error deleting node: %v\n", err)
	}

	list.printNodes()

	err = list.deleteFromIndex(1)
	if err != nil {
		fmt.Printf("Error deleting node: %v\n", err)
	}

	list.printNodes()

	err = list.deleteFromIndex(0)
	if err != nil {
		fmt.Printf("Error deleting node: %v\n", err)
	}

	list.printNodes()
}
