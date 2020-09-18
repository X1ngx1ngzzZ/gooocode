package main

import "math"

//2020715
//I
//设置一个全局的变量
var max int
func maxDepth(root *TreeNode) int {
	max=0
	if root ==nil{
		return max
	}
	deal(root,1)
	return max
}
//deal用于变更max的值
//让左右两子树去递归变更max的值
func deal(root *TreeNode,res int){
	if root!=nil {
		if max < res {
			max = res
		}

		deal(root.Left, res+1)

		deal(root.Right, res+1)

	}
}


//I

func isBalanced(root *TreeNode) bool {
	return dealBalance(root)!=-1
}

//左右子树为nil就返回0
//当正常情况时会返回自己左右子树的最高高度
//当左右的高度差大于1的时候会返回-1
func dealBalance(root *TreeNode)int{
	if root==nil{
		return 0
	}
	l:=dealBalance(root.Left)
	r:=dealBalance(root.Right)
	if l!=-1&&r!=-1&&int(math.Abs(float64(l-r))) <= 1{
		return int(math.Max(float64(l), float64(r))) + 1
	}

	return -1
}

//20200811
var depth int
func maxDepth(root *TreeNode) int {
	depth=0
	if root==nil{
		return depth
	}
	deal(root,1)
	return depth
}
func deal(treeNode *TreeNode,res int){
	if root!=nil{
		if max<res{
			depth=res
		}

	deal(root.Left,res+1)
	deal(root.Right,res+1)
	}
}