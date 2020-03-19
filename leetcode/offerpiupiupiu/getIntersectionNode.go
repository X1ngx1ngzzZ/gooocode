func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1，l2:=headA,headB
	m:=make(map[int]int)
	for i:=1;;i++{
		if l1==nil||l2==nil{
			return nil
		}
		if _,ok:=m[l2.Val];ok==true{
			return l2
		}
		if l1.Val==l2.Val{
			return l1
		}
		m[l1.Val]=i
		l1=l1.Next
		l2=l2.
		/*
		if value,ok:=m[l2.Val];ok==true{
			return l2
		}

		*/
	}
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1,l2:=headA,headB
	for l1!=l2{
	if l1==nil{
		l1=headB
	}else{
		l1=l1.Next
	}
	if l2==nil{
		l2=headA
	}else{
		l2=l2.Next
	}
	}
	return l1
	


}