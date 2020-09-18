package main

import (
	"fmt"
)


type TreeNode struct {
	Val string
	Left *TreeNode
	Right *TreeNode
}

func NewTree(value string) *TreeNode {
	return &TreeNode{
		Val : "a",
		Left : nil,//&TreeNode{},
		Right : nil,//&TreeNode{},
	}
}

// 增加左节点
func (node *TreeNode)addLeftNode(value string) {
	children := &TreeNode{
		Val : value,

	}
	node.Left = children
}

// 增加右节点
func (node *TreeNode)addRightNode(value string) {
	children := TreeNode{
		Val : value,
	}
	node.Right = &children
}


func main() {
	root := NewTree("a")
	root.Left.addLeftNode("b")

	fmt.Println(*root)
}