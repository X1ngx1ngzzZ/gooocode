


type TreeNode struct{
	Val int
	left *TreeNode
	right *TreeNode
}

func isSys(T *TreeNode)bool{
	if T==nil{
		return true
	}
	return isMor(T.left,T.right)
}

func isMor(p,q *TreeNode)bool{
	if p==nil||q==nil{
		return p==q
	}
	//没有比较值，要加p.vAL==q.Val
	return isMor(p.left,q.right)&&isMor(p.right,q.left)&&p.Val==q.Val
}

//20200612
func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	return isMor(root.Left,root.Right)
}

func isMor(p,q *TreeNode)bool{
	if p==nil||q==nil{
		return p==q
	}
	return isMor(p.Left,q.Right)&&isMor(p.Right,q.Left)&&p.Val==q.Val
}