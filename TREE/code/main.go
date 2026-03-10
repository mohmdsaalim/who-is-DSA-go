package main

import "fmt"

type Node struct {
	Val int
	Left *Node
	Right *Node
}

// insert into tree
func Insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Val: val}
	}
	if val < root.Val{
		root.Left = Insert(root.Left, val)
	}else{
		root.Right = Insert(root.Right, val)
	}
	return root
}
// in order 
func Inorder(root *Node)  {
	if root == nil{
		return
	}
	Inorder(root.Left)
	fmt.Print(root.Val, "-> ")
	Inorder(root.Right)
}
// post order
func Postorder(root *Node)  {
	if root == nil{
		return
	}
	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Print(root.Val, "-> ")
}
// pre order
func preorder(root *Node)  {
	if root == nil{
		return
	}
	fmt.Print(root.Val, "-> ")
	Postorder(root.Left)
	Postorder(root.Right)

}

func main() {

	var root *Node

	root = Insert(root, 10)
	root = Insert(root, 1)
	root = Insert(root, 2)
	root = Insert(root, 20)
	root = Insert(root, 30)
	root = Insert(root, 3)
	root = Insert(root, 4)


	fmt.Println("Inorder")
	Inorder(root)

	fmt.Println("post order")
	Postorder(root)

	fmt.Println("pre order")
	preorder(root)
}


func In(root *Node)  {
	if root == nil{
		return
	}
	In(root.Left) // l 
	fmt.Print(root.Val, "->") // root
	In(root.Right) // right
	
	// root-> left-> right pre
	// post left-> right -> root

}

//    complted 3leetcode based on traversal
