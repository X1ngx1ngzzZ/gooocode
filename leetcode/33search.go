package main

import "fmt"

//分有序无序
[4,5,6,7,0,1,2,3]
func search(nums []int, target int) int {
	left,right:=0,len(nums)-1

	for left<right{
	mid:=(left+right)/2
	if nums[mid]==target{
		return mid
	}
	if nums[left]<nums[mid]{
		if target>=nums[left]&&target<nums[mid]{
			right=mid-1
		}else{
			right=mid+1
		}
	}else{
		if target>nums[mid]&&target<=nums[right]{
			left=mid+1
		}else{
			right=mid-1
		}
	}
}
}

//找到翻转点，两个数组进行二分
/*
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	min := findMin(nums)
	fmt.Println(min)
	//说明在左半边
	if target < nums[0] || min == 0 {
		if serchdeal(nums, target, min, len(nums)) == -1 {
			return -1
		} else {
			return serchdeal(nums, target, min, len(nums))
		}
	} else {
		if serchdeal(nums, target, 0, min-1) == -1 {
			return -1
		} else {
			return serchdeal(nums, target, 0, min-1)
		}

	}

}
*/
func serchdeal(nums []int, target int, left, right int) int {
	//	left, right := 0, len(nums)-1
	for left <= right {
		//	fmt.Println("here")
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return -1
}
[4,5,6,7,0,1,2,3]
func findMin(nums []int) int {

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid-1
		} else {
			left = mid + 1
		}
	}
	return left
}



func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	min := findMin(nums)
	left, right := 0, 0
	//数组里只有两个元素
	//右边
	if nums[0] > target || min == 0 {
		left, right = min, len(nums)-1
	} else {
		left, right = 0, min-1
	}
	return serchdeal(nums, target, left, right)
}
