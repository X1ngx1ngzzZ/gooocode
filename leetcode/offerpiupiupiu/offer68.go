package main


//2020716
//I
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val>q.Val{
		p,q=q,p
	}
	if root==nil||p.Val<=root.Val&&q.Val>=root.Val{
		return root
	}else if p.Val>=root.Val{
		return lowestCommonAncestor(root.Right,p,q)
	}else {
		return lowestCommonAncestor(root.Left,p,q)
	}
}

//II
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==nil||p.Val==root.Val||q.Val==root.Val{
		return root
	}
	l:=lowestCommonAncestor(root.Left,p,q)
	r:=lowestCommonAncestor(root.Right,p,q)

	if l==nil{
		return r
	}else if r==nil{
		return l
	}else {
		return root
	}
}