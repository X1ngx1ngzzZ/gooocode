package main

import (
	"fmt"
)

func main() {
	m := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxPro(m))
}

func maxPro(arry []int) int {
	//特殊情况
	if arry == nil {
		return -1
	}
	result := 0
	//最开始的价格要很大
	minPrice := 1 << 31
	for _, v := range arry {
		//存储当前最小的
		minPrice = min(minPrice, v)
		//把当前的卖家减去之前存储的最小的
		result = max(result, v-minPrice)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
