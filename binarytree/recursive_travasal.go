package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

/**
*           1
		   / \
		  2   3
		/   /   \
	   4   5     7
	        \
			 6
*/
func main() {
	root := &Node{val: 1}
	first := &Node{val: 2}
	second := &Node{val: 3}
	third := &Node{val: 4}
	forth := &Node{val: 5}
	five := &Node{val: 6}
	six := &Node{val: 7}

	// build binary tree
	root.left = first
	root.right = second
	first.left = third
	second.left = forth
	second.right = six
	forth.right = five

	fmt.Println("preorder:")
	recursivePreOrderTravasal(root)
	fmt.Println("\ninorder:")
	recursiveInOrderTravasal(root)
	fmt.Println("\npostorder:")
	recursivePostOrderTravasal(root)
}

// 前序遍历
func recursivePreOrderTravasal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.val)
	recursivePreOrderTravasal(root.left)
	recursivePreOrderTravasal(root.right)
}

// 中序遍历
func recursiveInOrderTravasal(root *Node) {
	if root == nil {
		return
	}
	recursiveInOrderTravasal(root.left)
	fmt.Printf("%d ", root.val)
	recursiveInOrderTravasal(root.right)
}

// 后续遍历
func recursivePostOrderTravasal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.val)
	recursivePostOrderTravasal(root.left)
	recursivePostOrderTravasal(root.right)
}
