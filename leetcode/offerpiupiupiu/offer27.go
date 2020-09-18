package main


func mirrorTree(root *TreeNode) *TreeNode {
	return change(root)
}

func change(root *TreeNode) *TreeNode{
	if root==nil{
		return nil
	}
	root.Left,root.Right=root.Right,root.Left
	change(root.Left)
	change(root.Right)
	return root
}

//20200811
func mirrorTree(root *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	root.Left,root.Right=root.Right,root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}