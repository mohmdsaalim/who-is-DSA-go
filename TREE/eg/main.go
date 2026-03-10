package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func Insert(root *Node, val int) *Node {

	if root == nil {
		return &Node{Val: val}
	}

	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}

	return root
}

func Search(root *Node, key int) bool {

	if root == nil {
		return false
	}

	if root.Val == key {
		return true
	}

	if key < root.Val {
		return Search(root.Left, key)
	}

	return Search(root.Right, key)
}

func findMin(root *Node) *Node {

	for root.Left != nil {
		root = root.Left
	}

	return root
}

func Delete(root *Node, key int) *Node {

	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = Delete(root.Left, key)

	} else if key > root.Val {
		root.Right = Delete(root.Right, key)

	} else {

		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		temp := findMin(root.Right)
		root.Val = temp.Val
		root.Right = Delete(root.Right, temp.Val)
	}

	return root
}

func Inorder(root *Node) {

	if root == nil {
		return
	}

	Inorder(root.Left)
	fmt.Print(root.Val, " ")
	Inorder(root.Right)
}

func main() {

	var root *Node

	root = Insert(root, 10)
	root = Insert(root, 5)
	root = Insert(root, 15)
	root = Insert(root, 2)
	root = Insert(root, 7)

	fmt.Println("Inorder:")
	Inorder(root)

	fmt.Println("\nSearch 7:", Search(root, 7))

	root = Delete(root, 5)

	fmt.Println("\nAfter Delete:")
	Inorder(root)
}