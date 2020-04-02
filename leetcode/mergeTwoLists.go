//合并两个有序链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1==nil{
        return l2
    }
    if l2==nil{
        return l1
    }
    if l1.Val<=l2.Val{
        l1.Next=mergeTwoLists(l1.Next,l2)
        return l1
    }else{
        l2.Next=mergeTwoLists(l1,l2.Next)
        return l2
    }
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1==nil{
        return l2
    }
    if l2==nil{
        return l1
    }
    if l1.Val<l2.Val{
        l1.Next=mergeTwoLists(l1.Next,l2)
        return l1
    }else{
        l2.Next=mergeTwoLists(l1,l2.Next)
        return l2
    }
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1==nil{
        return l2
    }
    if l2==nil{
        return l1
    }
    if l1.Val<l2.Val{
        l1.Next=mergeTwoLists(l1.Next,l2)
        //这里记得返回值
        return l1
    }else{
        l2.Next=mergeTwoLists(l2.Next,l1)
        return l2
    }
}