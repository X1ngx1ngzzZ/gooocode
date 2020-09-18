package main

func missingNumber(nums []int) int {
	i:=0
	j:=len(nums)
	for i<j{
		tmp:=(i+j)/2
		if nums[tmp]==tmp{
			i=tmp+1
		}else {
			j=tmp
		}
	}
	return i
}