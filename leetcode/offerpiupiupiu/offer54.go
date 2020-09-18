package main


func kthLargest(root *TreeNode, k int) int {
	//二叉搜索树的中序遍历为 递增序列
	var res []int
	res=getnumbs(root,&res)
	return res[k-1]
}

func getnumbs(root  *TreeNode,res *[]int)[]int{
	if root.Right!=nil{
		getnumbs(root.Right,res)
	}
	if root!=nil{
		*res=append(*res,root.Val)
	}
	if root.Left!=nil{
		getnumbs(root.Left,res)
	}
	return *res
}