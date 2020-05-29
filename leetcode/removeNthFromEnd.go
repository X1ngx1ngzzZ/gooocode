package main

import "fmt"

/*
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l:=func(head)
    l1:=l+1-n
	l2:= fanhui(head,n)
	return l2

}
func len(head *ListNode)int{
	 pre :=head
	 count:=0

	for pre!= nil{
		count++
		pre = pre.Next

	}
	return count
}

func fanhui(head *ListNode,l int)int{
pre:=head
for i:=0;i<l;i++{
	pre=pre.Next
}
pre.Next=pre.Next.Next

}
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := new(ListNode)
	pre.Next = head
	fast, slow := head, head
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
	//return head
}

/*
fast,slow:=head,head
    for i:=0;i<=n;i++{
        fast=fast.Next
    }
    for fast!=nil{
        fast=fast.Next
        slow=slow.Next
    }
    slow.Next=slow.Next.Next
	return head
*/
/*
	type ListNode struct {
		Val int
		Next *ListNode
	}

   func main() {
	head := &ListNode{}
	head.Val = 1
	fmt.Println(head.Val)

   }


func main() {
head := new(ListNode)
head.Val = 1
fmt.Println(head.Val)

}
type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
head := ListNode{}
head.Val = 1
fmt.Printf("%T",head)

}
*/


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre:=make(ListNode)
	pre.Next=head
	fast,slow:=pre
	for i:=0;i<=n:i++{
		fast=fast.Next
	}
	for fast!=nil{
		fast=fast.Next
		slow=slow.Next
	}
	slow.Next=slow.Next.Next
	//
	return pre.Next
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

func main() {
	pre := new(*ListNode)
	fmt.Println(pre)
	var cur *ListNode
	fmt.Println(cur)
}
