type ListNode struct {
	    Val int
	     Next *ListNode
	 }
	
/*
func reverseList(head *ListNode) *ListNode {
	l:=new(listNode)
	pre:=new(listNode)
	pre=head
	for pre.Next!=nil{
		//l.val=pre.Val
		l2:=new(listNode)
		l2.Val=pre.Val
		l=l2
		l2.Next=
		pre=pre.Next
	}
}
*/

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
    pre=nil
	cur:=head
	for cur!=nil{
		temp:=cur.Next
		cur.Next=pre
		pre=cur
		cur=temp
	}
	return pre
}

func reverseList(head *ListNode) *ListNode {
    var last *ListNode
    pre:=head
    for pre!=nil{
        pre.Next,last,pre = last,pre,pre.Next
    }
    return last
}


func reverseList(head *ListNode) *ListNode {
 last:=new(listNode)
 last=nil
 pre:=head
 for pre!=nil{
	temp:=pre.Next
	pre.Next=last
	last=pre
	pre=temp
 }
 return last

	
}