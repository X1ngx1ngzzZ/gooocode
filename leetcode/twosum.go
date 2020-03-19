package main

import (
	"fmt"
)

//暴力解
/*
func twoSum(nums []int, target int) []int {
	s1 := []int{}
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if k2 > k1 {
				if target == v1+v2 {
					s1 = []int{k1, k2}
					break
				}
			}
		}
	}
	return s1
}
*/

// 使用Map

/*
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	res := []int{}
	for k, v := range nums {
//		m[v]=k
		if _, val2 := m[target-v]; val2 {
			res = []int{k, m[target-v]}
			break
		}
		m[v] = k
	}
	return res
}
*/
//自己使用Map

//把数组转换成Map
/*
func trans(nums []int) map[int]int {
	m := make(map[int]int)
	//var m1 map[int]int
	for k, v := range nums {
		m[k] = v
	}
	return m
}
*/

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	res := make([]int, 2, 2)

	for k, v := range nums {
		if _, v1 := m[target-v]; v1 {
			res = []int{m[target-v], k}
			break
		}
		m[v] = k
	}
	return res

}

func main() {
	ss1 := []int{5, 7, 6, 15}
	s := 11
	//	fmt.Println(trans(ss1))
	fmt.Println(twoSum(ss1, s))
}



func twoSum(nums []int, target int) []int {
	m:=make(map[int]int)
	re:=make([]int,2,2)
	for k,v:=range nums{
		temp:=target-nums[k]
		
		if num,ok:=m[temp]{
			return re[num,k]
		}
		m[v]=k	
	}
}