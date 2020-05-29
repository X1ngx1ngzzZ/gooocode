 type ListNode struct {
	     Val int
	     Next *ListNode
	 }

	 
//辅助栈解决
 func isPalindrome(head *ListNode) bool {
	if head==nil{
		return true
	}
	s:=make([]int, 0)
		pre:=head
		//入栈
		for pre!=nil{
			s=append(s,pre.Val)
			pre=pre.Next
		}	
		//比较
		for i:=0;i<len(s);i++{
			if s[len(s)-1]!=head.Val{
				return false
			}
			head=head.Next
			s=s[:len(s)-1]
		}
		return true
	   }



//反转链表然后对比
//这做法是有问题的，除非先生成一条新链表
func isPalindrome(head *ListNode) bool {
    tmp:=reverseList(head)
    for tmp!=nil{
        if tmp.Val!=head.Val{
            return false
        }
        tmp=tmp.Next
        head=head.Next
    }
    return true
}

func reverseList(head *ListNode) *ListNode {
	cur := new(ListNode)
	cur = nil
	//	var cur *ListNode
	cur = nil
	pre := head
	for pre != nil {
		pre, cur, pre.Next = pre.Next, pre, cur
	}
	return cur
}