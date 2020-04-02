//链表中倒数第K个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	pre:=new(ListNode)
	pre.Next=head
	fast,slow:=pre,pre
	for i:=0;i<=k;i++{
		fast=fast.Next
	}
	for fast!=nil{
		fast=fast.Next
		slow=slow.Next
	}
	return slow.Next
}


func getKthFromEnd(head *ListNode, k int) *ListNode {
	//特殊情况：传入List为nil，k的值大于链表长度
	if head ==nil{
		return nil
	}
	fast,slow:=head,head
	for i:=0;i<k;i++{
		fast=fast.Next
		if fast ==nil {
			return nil
		}
	}
	for fast!=nil{
		fast=fast.Next
		slow=slow.Next
	}
	return slow
}
func getKthFromEnd(head *ListNode, k int) *ListNode {
	//特殊情况：传入List为nil，k的值大于链表长度
	if head ==nil{
		return nil
	}
	fast,slow:=head.Next,head.Next
	for i:=0;i<k;i++{
		fast=fast.Next
		if fast ==nil {
			return nil
		}
	}
	for fast!=nil{
		fast=fast.Next
		slow=slow.Next
	}
	return slow.Next
}