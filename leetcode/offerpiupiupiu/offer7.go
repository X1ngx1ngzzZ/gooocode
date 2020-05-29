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
	//这个题里这个地方很关键，递归结束的终点，要重视啊
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
		//这里一定要注意数组的溢出,[:]中输出的是不包含后面的元素的
		Left:  buildTree(preorder[1:root+1], inorder[0:root]),
		right: buildTree(preorder[root+1:], inorder[root+1:]),
	}

}

