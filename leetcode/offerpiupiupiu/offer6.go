package main

type ListNode struct {
	Value int
	Next  *ListNode
}

func BuildList() {

}

//解法1，用切片代替了栈，遍历链表把当前值作为倒序输出
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	//其实这里可以省略，切片是自增长的，直接[]int{}
	//看链表长度，这里的看链表长度是可以倒序放进去
	l := LongOfList(head)
	//建个切片
	s := make([]int, l)
	now := new(ListNode)
	//把头指针的下一个给now，也就是第一个节点
	now = head
	//
	for i := l - 1; now != nil && i >= 0; {
		s[i] = now.Value
		now = now.Next
		i--
	}
	return s

}

//遍历链表看长度
func LongOfList(head *ListNode) int {
	long := 0
	pre := new(ListNode)
	pre = head
	for pre != nil {
		pre = pre.Next
		long++
	}
	return long
}

//解法2递归
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return appendData(head)
}

func appendData(head *ListNode) []int {
	if head.Next != nil {
		//如果要先输出这个节点的值，那么就先输出它后面的节点
		list := appendData(head.Next)
		//把当前节点的值加入到切片里
		list = append(list, head.Val)
		//然后返回切片
		return list
	}
	//返回以当前值产生的切片
	return []int{head.Val}
}
