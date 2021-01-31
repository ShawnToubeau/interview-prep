package main

import (
	"errors"
	"fmt"
	"strconv"
)

type BTNode struct {
	value      int
	leftChild  *BTNode
	rightChild *BTNode
}

type BinaryTree struct {
	root *BTNode
	size int
}

// Traverses a binary tree in order.
func (tree *BinaryTree) traverseInOrder(node *BTNode) {
	if node != nil {
		tree.traverseInOrder(node.leftChild)
		fmt.Printf("%v\n", node.value)
		tree.traverseInOrder(node.rightChild)
	}
}

// Traverses a binary tree in pre-order.
func (tree *BinaryTree) traverseInPreOrder(node *BTNode) {
	if node != nil {
		fmt.Printf("%v\n", node.value)
		tree.traverseInPreOrder(node.leftChild)
		tree.traverseInPreOrder(node.rightChild)
	}
}

// Traverses a binary tree in post-order.
func (tree *BinaryTree) traverseInPostOrder(node *BTNode) {
	if node != nil {
		tree.traverseInPostOrder(node.leftChild)
		tree.traverseInPostOrder(node.rightChild)
		fmt.Printf("%v\n", node.value)
	}
}

// Inserts a new node into the tree.
func (tree *BinaryTree) insert(value int) {
	newNode := BTNode{value: value}
	// set the root node if nil
	if tree.root == nil {
		tree.root = &newNode
		tree.size++
	} else { // else loop through children and insert
		current := tree.root

		// while our current node isn't nil
		for current != nil {
			// check if we want to insert on the left
			if newNode.value < current.value {
				// set left child if nil
				if current.leftChild == nil {
					current.leftChild = &newNode
					break
				}
				// continue down
				current = current.leftChild
			} else if newNode.value > current.value { // else, insert on right (change to >= to allow dupes)
				// set right child if nil
				if current.rightChild == nil {
					current.rightChild = &newNode
					break
				}
				// continue down
				current = current.rightChild
			} else {
				break
			}
		}
	}
}

// Checks whether the tree contains a specific value.
func (tree *BinaryTree) contains(value int) bool {
	current := tree.root
	// while the current node isn't nil
	for current != nil {
		// if the current node matches
		if current.value == value {
			return true
		}

		// traverse left side of tree
		if value < current.value {
			current = current.leftChild
		} else if value > current.value { // traverse right side of tree
			current = current.rightChild
		}
	}

	return false
}

// Retrieves the minimum value from the tree.
func (tree *BinaryTree) getMin() (int, error) {
	current := tree.root
	// if the tree is empty throw an error
	if current == nil {
		return 0, errors.New("tree is empty")
	}
	// traverse to the left-most child
	for current.leftChild != nil {
		current = current.leftChild
	}

	return current.value, nil
}

// Retrieves the maximum value from the tree.
func (tree *BinaryTree) getMax() (int, error) {
	current := tree.root
	// if the tree is empty throw an error
	if current == nil {
		return 0, errors.New("tree is empty")
	}
	// traverse to the left-most child
	for current.rightChild != nil {
		current = current.rightChild
	}

	return current.value, nil

}

// Deletes the minimum value within the tree.
func (tree *BinaryTree) deleteMin() error {
	current := tree.root
	prev := &BTNode{}
	// if the tree is empty throw an error
	if current == nil {
		return errors.New("tree is empty")
	}
	// traverse to the left-most child
	for current.leftChild != nil {
		prev = current
		current = current.leftChild
	}

	// set the left-most node to nil
	prev.leftChild = nil

	return nil
}

// Deletes the maximum value within the tree.
func (tree *BinaryTree) deleteMax() error {
	current := tree.root
	prev := &BTNode{}
	// if the tree is empty throw an error
	if current == nil {
		return errors.New("tree is empty")
	}
	// traverse to the right-most child
	for current.rightChild != nil {
		prev = current
		current = current.rightChild
	}

	// set the right-most node to nil
	prev.rightChild = nil

	return nil
}

// Deletes a specific node from the tree.
func (tree *BinaryTree) delete(value int) error {
	delNode := tree.root
	delNodeParent := &BTNode{}
	// whether the node we're deleting is a right child or not
	isRightChild := false

	for delNode.value != value && delNode != nil {
		delNodeParent = delNode

		if value < delNode.value {
			delNode = delNode.leftChild
			isRightChild = false
		} else if value > delNode.value {
			delNode = delNode.rightChild
			isRightChild = true
		}
	}

	if delNode == nil {
		return errors.New("value not found")
	}

	// node is a leaf
	if delNode.leftChild == nil && delNode.rightChild == nil {
		tree.handleNodeSwap(delNodeParent, nil, isRightChild)
	} else if delNode.leftChild != nil && delNode.rightChild == nil { // node has left child
		tree.handleNodeSwap(delNodeParent, delNode.leftChild, isRightChild)
	} else if delNode.leftChild == nil && delNode.rightChild != nil { // node has right child
		tree.handleNodeSwap(delNodeParent, delNode.rightChild, isRightChild)
	} else if delNode.leftChild != nil && delNode.rightChild != nil { // node has left and right children
		// node to swap with the deleted node
		replacement := delNode.rightChild
		// we'll need to know the replacement's parent to remove their link as well
		replacementParent := &BTNode{}
		// find the smallest element under the right child of the node we're deleting
		for replacement.leftChild != nil {
			replacementParent = replacement
			replacement = replacement.leftChild
		}
		// delete old connection to the replacement node
		if replacementParent != nil {
			replacementParent.leftChild = nil
		}
		// set the children of the replacement to the children of our node to delete
		replacement.leftChild, replacement.rightChild = delNode.leftChild, delNode.rightChild

		tree.handleNodeSwap(delNodeParent, replacement, isRightChild)
	}

	return nil
}

// Helper function for delete. Replaces the node we're deleting with the node that's taking it's place.
func (tree *BinaryTree) handleNodeSwap(delNodeParent *BTNode, replacement *BTNode, isRightChild bool) {
	// if the parent of our node to delete is null we are deleting the root
	if *delNodeParent == (BTNode{}) {
		tree.root = replacement
	} else if isRightChild {
		delNodeParent.rightChild = replacement
	} else {
		delNodeParent.leftChild = replacement
	}
}

func (tree *BinaryTree) printTree() {
	if tree.root == nil {
		fmt.Println("Tree is empty")
		return
	}

	str := strconv.Itoa(tree.root.value)

	pointerRight := "└──"
	pointerLeft := "└──"

	if tree.root.rightChild != nil {
		pointerLeft = "├──"
	}

	tree.printChildren(&str, "", pointerLeft, tree.root.leftChild, tree.root.rightChild != nil)
	tree.printChildren(&str, "", pointerRight, tree.root.rightChild, false)

	fmt.Printf("%v\n", str)
}

func (tree *BinaryTree) printChildren(str *string, padding string, pointer string, node *BTNode, hasRightSibling bool) {
	if node != nil {
		*str += "\n" + padding + pointer + strconv.Itoa(node.value)

		paddingForBoth := padding

		if hasRightSibling {
			paddingForBoth += "│  "
		} else {
			paddingForBoth += "   "
		}

		pointerRight := "└──"
		pointerLeft := "└──"

		if node.rightChild != nil {
			pointerLeft = "├──"
		}

		tree.printChildren(str, paddingForBoth, pointerLeft, node.leftChild, node.rightChild != nil)
		tree.printChildren(str, paddingForBoth, pointerRight, node.rightChild, false)
	}
}

func main() {
	tree := BinaryTree{}

	tree.insert(40)
	tree.insert(20)
	tree.insert(9)
	tree.insert(32)
	tree.insert(15)
	tree.insert(8)
	tree.insert(27)
	tree.insert(55)
	tree.printTree()

	contains := tree.contains(22)
	fmt.Printf("Contains: %v\n", contains)
	contains = tree.contains(20)
	fmt.Printf("Contains: %v\n", contains)
	contains = tree.contains(55)
	fmt.Printf("Contains: %v\n", contains)

	min, err := tree.getMin()
	if err != nil {
		fmt.Printf("Error getting minimum: %v\n", err)
	} else {
		fmt.Printf("Min: %v\n", min)
	}

	max, err := tree.getMax()
	if err != nil {
		fmt.Printf("Error getting maximum: %v\n", err)
	} else {
		fmt.Printf("Max: %v\n", max)
	}

	err = tree.deleteMin()
	if err != nil {
		fmt.Printf("Error deleting minimum value: %v\n", err)
	}

	err = tree.deleteMax()
	if err != nil {
		fmt.Printf("Error deleting maximum value: %v\n", err)
	}

	tree.printTree()

	err = tree.delete(40)
	if err != nil {
		fmt.Printf("Error deleting value: %v\n", err)
	}

	tree.printTree()

	err = tree.delete(9)
	if err != nil {
		fmt.Printf("Error deleting value: %v\n", err)
	}

	tree.printTree()

	//fmt.Println("In order..")
	//tree.traverseInOrder(tree.root)
	//fmt.Println("Pre order..")
	//tree.traverseInPreOrder(tree.root)
	//fmt.Println("Post order..")
	//tree.traverseInPostOrder(tree.root)
}
