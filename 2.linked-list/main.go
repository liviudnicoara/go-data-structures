package main

import "fmt"

// Problem Statement:
// Reverse a linked list. Do it in-place.

// Example:
// Input: 1 -> 2 -> 3 -> 4 -> 5
// Output: 5 -> 4 -> 3 -> 2 -> 1

type Node struct {
	val  int
	next *Node
}

func printList(head *Node) {
	for head != nil {
		fmt.Print(head.val, "->")
		head = head.next
	}

	fmt.Println("nil")
}

func reverseList(head *Node) *Node {
	var prev *Node

	for head != nil {
		nextNode := head.next
		head.next = prev
		prev = head
		head = nextNode
	}

	return prev
}

func main() {
	// 1 -> 2 -> 3 -> 4 -> 5
	node := &Node{val: 1, next: &Node{val: 2, next: &Node{val: 3, next: &Node{val: 4, next: &Node{val: 5}}}}}
	printList(node)

	// 5 -> 4 -> 3 -> 2 -> 1
	reverse := reverseList(node)
	printList(reverse)
}
