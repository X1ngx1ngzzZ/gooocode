package main

import "golang.org/x/text/unicode/bidi"

//重建二叉树
//给出深度的前序和中序，建立二叉树

func main() {

}

type TreeNode{
	Val int
	Left *TreeNode
	right *TreeNode
}


func buildTree(preorder []int, inorder []int) *TreeNode {
	//这个题里这个地方很关键，递归结束的终点，要重视啊
	if preorder == nil || inorder == nil {
		return nil
	}

	var root int
	//找出前序遍历的第一个在中序遍历中的位置
	for k, v := range inorder {
		if preorder[0] == v {
			root = k
			//退的记得出循环
			break
		}
	}

	//Treenode.left:preorder:preorder[1:root]    inorder:inorder[:k-1]
	return &TreeNode{
		Val: preorder[0],
		//这里一定要注意数组的溢出,[:]中输出的是不包含后面的元素的
		Left:  buildTree(preorder[1:root+1], inorder[0:root]),
		right: buildTree(preorder[root+1:], inorder[root+1:]),
	}

}

//2020.7.12
func buildTree(preorder []int, inorder []int) *TreeNode {
	//停止条件,需要加判断前后的长度为0
	if preorder==nil||inorder==nil||len(preorder)==0||len(inorder)==0{
		return nil
	}
	var root int
	//遍历找位置
	for k,v:=range inorder{
		if preorder[0]==v{
			root=k
			break
		}
	}

	return &TreeNode{
		Val: preorder[0],
		Left:buildTree(preorder[1:root+1],inorder[:root]),
		Right:buildTree(preorder[root+1:],inorder[root+1:]),
	}

}

//20200811
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0||len(inorder)==0{
		return nil
	}
	var root int
	for k,v:=range inorder{
		if v==preorder[0]{
			root=k
			break
		}
	}
	return &TreeNode{
		Val:preorder[0],
		left:buildTree(preorder[1:root+1],inorder[:root]),
		right:buildTree(preorder[root+1:],inorder[root+1:]),
	}
}

//20200816
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0||len(inorder)==0{
		return nil
	}
	var root int
	for k,v:=range inorder{
		if preorder[0]==v{
			root=k
			break
		}
	}
	return &TreeNode{
		Val:preorder[0],
		Left:buildTree(preorder[1:root+1],inorder[:root]),
		Right:buildTree(preorder[root+1:],inorder[root+1:]),
	}
}