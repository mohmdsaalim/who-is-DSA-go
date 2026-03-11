package main

import "fmt"

type Node struct {
	Val int
	left *Node
	right *Node
}

func insert(root *Node, val int) *Node {
	if root == nil{
		return &Node{Val: val}
	}
	if val < root.Val{
		root.left = insert(root.left, val)
	}else{
		root.right = insert(root.right, val)
	}

	return root
}


func Search(root *Node, val int) bool {
	if root == nil{
		return false
	}
	if root.Val == val{
		return true
	}
	if val < root.Val{
		return Search(root.left, val)
	}
	return Search(root.right, val)
}


func Inorder(root *Node)  {
	if root == nil{
		return
	}
	Inorder(root.left)
	fmt.Print(root.Val, "-")
	Inorder(root.right)
}

func preorder(root *Node)  {
	if root == nil{
		return
	}
	fmt.Print(root.Val, "-")
	Inorder(root.left)
	Inorder(root.right)
}

func postorder(root *Node)  {
	if root == nil{
		return
	}
	Inorder(root.left)
	Inorder(root.right)
	fmt.Print(root.Val, "-")
}


func main() {
	var root *Node
	root = insert(root, 10)
	root = insert(root, 1)
	root = insert(root, 2)
	root = insert(root, 3)
	root = insert(root, 11)
	root = insert(root, 12)
	root = insert(root, 13)
	

	fmt.Println("helo")
	Inorder(root)
	a := Search(root, 1)
	fmt.Println(a)
	// postorder(root)
	// preorder(root)
}