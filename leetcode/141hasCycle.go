type ListNode struct {
	    Val int
	     Next *ListNode
	 }
	
//快慢指针
func hasCycle(head *ListNode) bool {
	if head==nil{
		return false
	}
	pre:=head.Next
	fast:=pre
	slow:=pre
	//不是环的话循环结束条件
	for  slow!=nil&&fast!=nil&&fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow{
			return true
		}
	}
	return false
}

//用下map

func hasCycle(head *ListNode) bool {
	if head==nil{
		return false
	}
	m:=make(map[*ListNode]int)
	//遍历链表
	for head!=nil{
		if _,ok:=m[head];ok{
			return true
		}
		m[head]=head.Val
		head=head.Next
	}
	return false
}