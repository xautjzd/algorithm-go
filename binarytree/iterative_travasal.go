package main

import (
	"fmt"
	"sync"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Stack struct {
	mu sync.Mutex
	s  []*Node
}

func NewStack() *Stack {
	return &Stack{
		sync.Mutex{},
		make([]*Node, 0),
	}
}

func (s *Stack) Push(node *Node) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.s = append(s.s, node)
}

func (s *Stack) Peek() *Node {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.s[len(s.s)-1]
}

func (s *Stack) Pop() *Node {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.s) == 0 {
		return nil
	}
	node := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]

	return node
}

func (s *Stack) Empty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.s == nil || len(s.s) == 0 {
		return true
	}
	return false
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
	iterativePreOrderTravasal(root)
	fmt.Println("\ninorder:")
	iterativeInOrderTravasal(root)
	// fmt.Println("\ninorder:")
}

// 前序遍历
func iterativePreOrderTravasal(root *Node) {
	if root == nil {
		return
	}
	stack := NewStack()
	stack.Push(root)
	for !stack.Empty() {
		current := stack.Pop()
		fmt.Printf("%d ", current.val)
		if current.right != nil {
			stack.Push(current.right)
		}
		if current.left != nil {
			stack.Push(current.left)
		}
	}
}

// 中序遍历
func iterativeInOrderTravasal(root *Node) {
	if root == nil {
		return
	}
	stack := NewStack()
	current := root

	for !stack.Empty() || current != nil {
		if current != nil {
			stack.Push(current)
			current = current.left
		} else {
			current = stack.Pop()
			fmt.Printf("%d ", current.val)
			current = current.right
		}
	}
}

// TODO 未完成
func iterativePostOrderTravasal(root *Node) {
	if root == nil {
		return
	}
	stack := NewStack()
	current := root
	for !stack.Empty() || current != nil {
		if current != nil {
			if current.right != nil {
				stack.Push(current.right)
			}
			stack.Push(current)
			current = current.left
		} else {
			current = stack.Pop()
			fmt.Printf("%d ", current.val)
		}
	}
}
