package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	//把res传进去
	dfs(root, &res)
	return res
}
//res指向的地址
func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	//
	l := dfs(root.Left, res)
	r := dfs(root.Right, res)
	*res = max(*res, l+r)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	m := 0
	a := &m

	fmt.Println(a)
	//fmt.Println(b)
}
