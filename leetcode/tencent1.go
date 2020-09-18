package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Value int
	Next *ListNode
}

func main(){
	var n,k int
	arryRes:=make([]int,0)
	reader:=bufio.NewReader(os.Stdin)
	fmt.Scan(&n,&k)
	input,_:=reader.ReadString('\n')
	result:=strings.Fields(input)
	if len(result)!=n{
		return
	}
	arrayInt:=transferInt(result)
	rowList:=makeListNode(arrayInt)
	res:=deal(rowList,k)
	for res!= nil {
		//fmt.Println(res.Value)
		arryRes=append(arryRes,res.Value)
		res = res.Next
	}
	//fmt.Println(arryRes)
	s:=strings.Join(transferString(arryRes)," ")
	fmt.Println(s)


}



func transferInt(arryS []string)[]int{
	res:=make([]int,len(arryS))
	for i:=0;i<len(arryS);i++{
		res[i],_=strconv.Atoi(arryS[i])
	}
	return res
}


func transferString(arryS []int)[]string{
	res:=make([]string,len(arryS))
	for i:=0;i<len(arryS);i++{
		res[i]=strconv.Itoa(arryS[i])
	}
	return res
}




func makeListNode(nums []int) *ListNode {

	if len(nums) == 0{
		return nil
	}

	res := &ListNode{
		Value:nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Value:nums[i],}
		temp = temp.Next
	}

	return  res
}

func deal(list *ListNode,k int)*ListNode{

	if k==0{
		return list
	}
	if k==1{
		return list.Next
	}

	root:=list
	for i := 1; i < k-1; i++ {
		root = root.Next
		fmt.Println(root.Value)
	}
	root.Next=root.Next.Next


	return list
}
