package main

import "fmt"

// Problem Statement:
// Given a binary tree, find its maximum depth.
// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := max(node.Left)
	right := max(node.Right)

	if left > right {
		return 1 + left
	}

	return 1 + right
}

func main() {

	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(max(tree))
}
