func maxDepth(root *TreeNode) int {
    if root==nil{
		return nil
	}
	count:=0
	for{
	if root.Left!=nil{
		count++
		maxDepth(root.Left)
	}else if root.Right!=nil{
		count++
		maxDepth(root.Right)
	}else{
		return count
	}
}

}

func maxDepth(root *TreeNode) int {
if root == nil {
	return 0
}
left := maxDepth(root.Left)
right := maxDepth(root.Right)
return max(left, right) + 1
}

func max(a, b int) int {
if a > b {
	return a
}
return b
}

func maxDepth(root *TreeNode) int {
	left:=maxDepth(root.Left)
	Right:=maxDepth(root.Right)
	return max(left,Right)+1
}

func max(a,b)int{
	if a > b {
		return a
	}
	return b
}