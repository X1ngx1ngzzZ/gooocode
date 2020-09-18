package main

import "fmt"

func hammingWeight(num uint32) int {
	var res int
	res=0
	for num>0{
		if num&1==1{
			res++
		}
		tmp:=num>>1
		num=tmp
	}
	return res
}
func main(){
	a:=uint32(00000000000000000000000000001011)
	fmt.Println(hammingWeight(a))

}