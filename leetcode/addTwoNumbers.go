package main

import (
	"fmt"
)

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

}

func main() {
	l1 := list.New()
	l1.Init()
	l1.PushBack(0)
	l1.PushBack(1)
	e1 := l1.PushFront(3)
	for e1; e1 != nil; e1 = e1.Next() {
		fmt.Println(e1.Value)
	}
	//fmt.Println(l1.Back())
	//fmt.Println(l1.Len())
	//移动的是节点不是数值
	//l1.MoveToBack(0)
	//fmt.Println(l1.Back())
	//l2:=list.New()
	//l2.Init()
}



type ListNode struct {
		Val int
    	Next *ListNode
	}

func addTwoNumbers(l1 *ListNode, l2 *ListNode){
	l3:=new(*ListNode)
	//var now1  now2 *int
	//now1 = &l1
	//now2 = &l2
	carry:=0
	sum:=0
	for {
		if l1 !=nil{
			sum+=l1.val
			l1=l1.Next
		}
		if l2!=nil{
			sum+=l2.val
			l2=l2.Next
		}
		sum+=carry
		carry=sum/10
		l3.val:=sum%10
		l3.next:=new(*ListNode)
		if l1==nil&&l2==nil&&carry==0{
			break
		}

	}
	return l3

}

type ListNode struct {
	Val int
   Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
l3:=&ListNode{}
  carry:=0
  sum:=0
  for {
	  if l1 !=nil{
		  sum+=l1.Val
		  l1=l1.Next
	  }
	  if l2!=nil{
		  sum+=l2.Val
		  l2=l2.Next
	  }
	  sum+=carry
	  carry=sum/10
	  l3.Val=sum%10
	  l3.Next=l3
	  if l1==nil&&l2==nil&&carry==0{
		  break
	  }

  }
  return l3

}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//l3:=&ListNode{}
  result:=&ListNode{}
  l3:=result
  carry:=0
  sum:=0
  for {
	  if l1 !=nil{
		  sum+=l1.Val
		  l1=l1.Next
	  }
	  if l2!=nil{
		  sum+=l2.Val
		  l2=l2.Next
	  }
	  sum+=carry
	  carry=sum/10
	  l3.Val=sum%10
	  if l1==nil&&l2==nil&&carry==0{
		  break
	  }
	  l3.Next = &ListNode{}
	  l3=l3.Next
	  sum=0
  }
  return result

}

func main() {
	l1 := &ListNode{}
	node := l1
	node.Val = 2
	node.Next = &ListNode{}
	node.Next = node
	node.Val = 4
	node.Next = &ListNode{}
	node.Next = node
	node.Val = 3
	fmt.Println(l1)
	l2 := &ListNode{}
	node = l2
	node.Val = 5
	node.Next = &ListNode{}
	node.Next = node
	node.Val = 6
	node.Next = &ListNode{}
	node.Next = node
	node.Val = 4
	fmt.Println(l2)
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3:=new(ListNode)
	sum:=0
	carry:=0
	for{
		if l1!=nil{
			sum+=l1.Val
			l1=l1.Next
		}
		if l2!=nil{
			sum+=l2.Val
			l2=l2.Next
		}
		sum+=carry
		carry=sum/10
		l3.Val=sum%10
		if l1==nil&&l2==nil&&carry==0{
			break
		}
		l4:=new(ListNode)
		l3.Next=l4
		l3=l3.Next
		sum=0
	}
}