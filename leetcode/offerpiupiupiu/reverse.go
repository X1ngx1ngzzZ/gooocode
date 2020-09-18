//反转链表
type ListNode struct {
	Val  int
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
	pre = nil
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func reverseList(head *ListNode) *ListNode {
	var last *ListNode
	pre := head
	for pre != nil {
		pre.Next, last, pre = last, pre, pre.Next
	}
	return last
}

func reverseList(head *ListNode) *ListNode {
	last := new(listNode)
	last = nil
	pre := head
	for pre != nil {
		temp := pre.Next
		pre.Next = last
		last = pre
		pre = temp
	}
	return last

}

//最终版本，两个指针，依次前移
func reverseList(head *ListNode) *ListNode {
	//第一个让它指为nil
	pre := new(listNode)
	//如果不让它为空他会为0
	pre = nil
	cur := head
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
		//	pre,cur,cur.Next = cur,cur.Next,pre
	}
	//一定是返回前面的指针，因为这时候后面的已经指到nil了
	return cur
}

//2020605
func reverseList(head *ListNode) *ListNode {
	pre := new(ListNode)
	pre = nil
	cur := head
	for cur != nil {
		pre, cur.Next, cur = cur, pre, cur.Next
	}
	//必须返回后面的
	return cur
}