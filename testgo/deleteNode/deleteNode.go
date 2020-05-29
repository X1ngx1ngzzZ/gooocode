package Node


func DeleteNode(head *ListNode, val int) *ListNode {
	if val==head.Val{
		return nil
	}
	
	pre,cur:=head,head.Next
	for cur!=nil{
		if cur.Val==val{
			pre.Next=cur.Next
			break
		}
		cur=cur.Next
		pre=pre.Next
	}
	return head
}


