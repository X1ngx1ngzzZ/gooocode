//删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	//特殊情况，给的值就是头指针
	if val==head.Val{
		return head.Next
	}
	//双指针，pre指向第一个，cur指向第二个
	pre,cur:=head,head.Next
	//第二个不为空的情况下
	for cur!=nil{
		//他的值为给的值
		if cur.Val == val{
			//在给指针的时候一定要注意，就是大白话pre的后面应该是cur的后面
			pre.Next= cur.Next
			break
		}else{
			cur=cur.Next
			pre=pre.Next
		}
	}
	return head

}


func deleteNode(head *ListNode, val int) *ListNode {
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


