//两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//判断特殊情况
	if headA==nil&&headB==nil{
		return nil
	}else if headA==nil{
		return headB
	}else if headB==nil{
		return headA
	}
	//让两个指针指向两个链表
	A,B:=headA,headB
	//一样的时候就停
	//如果不相交就会一起走到nil，返回nil
	for A!=B{
		if A==nil{
			A=headB
		}else{
			A=A.Next
		}
		if B==nil{
			B=headA
		}else{
			B=B.Next
		}
	}
	return A



}