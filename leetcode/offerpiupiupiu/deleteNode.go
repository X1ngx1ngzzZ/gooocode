func deleteNode(head *ListNode, val int) *ListNode {
	pre:=head
	for{
		if pre.Val==val{
			return head.Next
		}
		if pre.Next.Val==val{
			pre.Next=pre.Next.Next
			return head
		}
		pre=pre.Next
	}
	return head
}

