package main

//2020820

//map
func twoSum(nums []int, target int) []int {
	m:=make(map[int]int)
	res:=[]int{}
	for _,v:=range nums{
		if _,v1:= m[target-v];v1{
			res=append(res,target-v)
			res=append(res,v)
			return res
		}
		m[v]++
	}
	return nil
}

//双指针

func twoSum(nums []int, target int) []int {
	i:=0
	j:=len(nums)-1
	res:=[]int{}
	for i<j{
		if target==nums[i]+nums[j]{
			res=append(res,nums[i])
			res=append(res,nums[j])
		}else if target<nums[i]+nums[j]{
			j--
		}else if target>nums[i]+nums[j]{
			i++
		}
	}
	return res
}