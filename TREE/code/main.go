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

func inorder(root *Node)  {
	if root == nil{
		return
	}
	inorder(root.left)
	fmt.Println(root.Val)
	inorder(root.right) 
}


func main() {
	
}