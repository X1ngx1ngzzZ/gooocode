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