package main

import "fmt"

func main(){
	var a int
	fmt.Scan(&a)
	var res []int
	if a<2178{
		fmt.Println(0)
		return
	}
	if a>1e7{
		return
	}
	for i:=2178;i<a;i++{
		res=deal(i,a,res)
	}
	len:=len(res)
	if len==0{
		fmt.Println(0)
	}else {
		fmt.Println(len/2)
		for k:=0;k<len;k+=2{
			fmt.Println(res[k],res[k+1])
		}
	}

}
func deal(i ,a int,res []int)[]int{
	fan:=reverse(i)
	tmp:=i*4
	if tmp==fan{
		res=append(res,i)
		res=append(res,tmp)
	}
	return res
}

func reverse(x int) int {
	y := 0
	for x!=0 {
		y = y*10 + x%10
		if !( -(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}