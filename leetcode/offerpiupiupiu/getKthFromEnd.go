//链表中倒数第K个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	pre := new(ListNode)
	pre.Next = head
	fast, slow := pre, pre
	for i := 0; i <= k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Next
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	//特殊情况：传入List为nil，k的值大于链表长度
	if head == nil {
		return nil
	}
	fast, slow := head, head
	//这样直接取头结点会导致特殊情况的判别情况影响到正常的，所以正确的在后面
	for i := 0; i < k; i++ {
		fast = fast.Next
		if fast == nil {
			return nil
		}
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
func getKthFromEnd(head *ListNode, k int) *ListNode {
	//特殊情况：传入List为nil，k的值大于链表长度
	if head == nil {
		return nil
	}
	//建个中间值，让它从头结点前面开始
	pre := new(ListNode)
	pre.Next = head
	fast, slow := pre, pre
	for i := 0; i < k; i++ {
		fast = fast.Next
		//避免k值大于链表长度
		if fast == nil {
			return nil
		}

	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Next
}

//20200605
//特殊情况，如只有一个节点，要取倒数第一个，所以要从第一个之前开始