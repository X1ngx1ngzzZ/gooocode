package main

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
	if preorder == nil || inorder == nil {
		return nil
	}

	var root int
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
		//这里一定要注意数组的溢出
		Left:  buildTree(preorder[1:root+1], inorder[0:root]),
		right: buildTree(preorder[root+1:], inorder[root+1:]),
	}

}

