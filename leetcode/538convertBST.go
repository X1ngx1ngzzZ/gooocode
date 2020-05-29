//栈
func convertBST(root *TreeNode) *TreeNode {
	node := root
	tempNum := 0
	var stack [100]*TreeNode
	var top = -1
	for node != nil || top != -1{
		if node != nil{  
			top++
			stack[top] = node
			node = node.Right //先走右边
		}else{
			top--  //再走左边
			tempNum += stack[top+1].Val  //1.node = stack[top+1].left 右边走到最后叶子节点时，top回溯到上一节点
			stack[top+1].Val = tempNum   //2.用上一节的下一个节点，也就是当前节点 左边是否还有节点
			node = stack[top+1].Left
		}
	}
	return root


}


//递归
var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	dfs(root)
	return root
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
	}
}

