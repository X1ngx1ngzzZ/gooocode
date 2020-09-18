type ListNode struct {
	Val  int
	Next *ListNode
}

//反向打印每个节点的值
//先正向打印，再反过来
func reversePrint(head *ListNode) []int {
	tmp := []int{}
	pre := head
	for pre != nil {
		tmp = append(tmp, pre.Val)
		pre = pre.Next
	}
	res := []int{}
	for len(tmp) != 0 {
		res = append(res, tmp[len(tmp)])
		tmp = tmp[:len(tmp)-1]
	}
	return res
}

//或者可以直接遍历链表长度，反着放进去
