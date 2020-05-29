 type ListNode struct {
	     Val int
	     Next *ListNode
	 }
func detectCycle(head *ListNode) *ListNode {
	//找交点,证明有环
	if head==nil{
		return nil
	}
	fast,slow:=head,head
	
	for fast!=nil&&fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
	//如果有环，把fast重新给到初始位置
		if fast==slow{
			fast=head
			for {
				//一定要在这里先判断，否则链表为圆环
				//就会直接跳到第二个
				if fast==slow{
					return fast
				}
				fast=fast.Next
				slow=slow.Next
				
			}
		}
	}

	return nil
}
