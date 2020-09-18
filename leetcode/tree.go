package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//前序遍历

func upSort(root *TreeNode){
	if root==nil{
		return
	}
	fmt.Println(root.Val)
	upSort(root.Left)
	upSort(root.Right)
}