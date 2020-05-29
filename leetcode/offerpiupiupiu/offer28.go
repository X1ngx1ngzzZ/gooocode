


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