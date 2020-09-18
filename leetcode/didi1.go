package main

import (
	"fmt"
	"strconv"
)

func main(){
	var n int
	fmt.Scan(&n)
	restmp:=deal(n)
	fmt.Println(len(restmp)/2)
	res:=dealArry(restmp)
	for i:=0;i<len(res);i++{
		fmt.Println(res[i])
	}



}
func deal(n int)[]int{

	res:=make([]int,0)
	for i:=1;i<=9;i++{
		for j:=0;j<=9;j++{
			for k:=0;k<=9;k++{
				tmp1:=i*100+j*10+k
				tmp2:=i*100+k*10+k
				if tmp1+tmp2==n{
					if i!=j&&j!=k&&i!=k {
						res = append(res, tmp1)
						res = append(res, tmp2)
					}
				}
			}
		}
	}
	return res
}
func dealArry(arry []int)[]string{
	res:=make([]string,0)
	for i:=0;i<len(arry);i+=2{
		tmp1:=strconv.Itoa(arry[i])
		tmp2:=strconv.Itoa(arry[i+1])
		tmp3:=tmp1+" "+tmp2
		res=append(res,tmp3)
	}
	return res
}