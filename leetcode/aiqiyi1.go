package main

import (
	"fmt"

)

func main(){
	var num int
	fmt.Scan(&num)
	//tmp:=muti(num)
	//string:=strconv.Itoa(tmp)
	res:=deal(num)
	fmt.Println(res)
}

//func muti(n int)int{
//	res:=1
//	for i:=1;i<=n;i++{
//		res=i*res
//	}
//	return res
//}
//
//func deal(n int)int{
//	var res int
//	for n>0{
//		tmp:=n%10
//		if tmp==0{
//			res++
//			n=n/10
//		}else {
//			break
//		}
//	}
//	return res
//}

func deal(input int)int{
	if input<5{
		return 0
	}else {
		return input/5+deal(input/5)
	}
}