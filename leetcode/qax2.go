package main

import "fmt"

//func house( person []int ) int {
//	// write code here
//	var res int
//	for i:=1;i<=len(person);i++{
//		res+=i
//	}
//	return res
//}

func main(){
	test:=[]int{3,2,4}
	fmt.Println(house(test))
}
func house( person []int ) int {
	// write code here
	res_zheng:=make([]int,0)
	res_fan:=make([]int,0)
	for i:=0;i<len(person);i++{
		res_zheng=append(res_zheng,1)
	}
	for i:=0;i<len(person);i++{
		res_fan=append(res_fan,1)
	}
	for i:=0;i<len(person)-1;i++{
		if person[i]>person[i+1]{
			res_zheng[i]++
		}
	}
	for i:=len(person)-1;i>1;i--{
		if person[i]>person[i-1]{
			res_fan[i]++
		}
	}
	tmp1:=0
	tmp2:=0
	for _,v:=range res_zheng{
		tmp1+=v
	}
	for _,v:=range res_fan{
		tmp2+=v
	}
	if tmp2>tmp1{
		return tmp1+1
	}else {
		return tmp2+1
	}
}


