func isBalanced(root *TreeNode) bool {
	_,b:=cal(root)
	return b
	}
	
	func max(a, b int) int {
		if a <= b {
		 return b
		}
		return a
	   }
	   
	 func abs(a int) int {
		if a >= 0 {
		 return a
		}
		return -a
	   }
	
	func cal(n *TreeNode)(int,bool){ 
		if n==nil{
			return 0,true
		}
		l1,v1:=cal(n.Left)
		l2,v2:=cal(n.Right)
		if !v1||!v2||abs(l1-l2)>1{
			return 0,false
		}
	return max(l1,l2)+1,true
	}