package main

import (
	"fmt"
)

/*
func maxSlidingWindow(nums []int, k int) []int {
	s:=make([]int,0,3)
	s=nums[0:]
}
*/

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	s := make([]int, 0, 3)
	s = nums[1:5]
	fmt.Println(s)
}
