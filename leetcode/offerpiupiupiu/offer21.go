package main

//2020720
//自定义排序的冒泡。好像时间超了
func exchange(nums []int) []int {
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if bigger(nums[i],nums[j]){
				nums[i],nums[j]=nums[j],nums[i]
			}
		}
	}
	return  nums
}

func bigger(a,b int)bool {
	if a%2==0&&b%2==1{
		return true
	}
	return  false
}

//双指针
func exchange(nums []int) []int {
	l:=0
	r:=len(nums)-1
	for l<r{
		if nums[l]%2==0&&nums[r]%2!=0{
			nums[l],nums[r]=nums[r],nums[l]
		}
		//这里可以用for判断，就可以不用每次都跳到外循环判断
		if nums[l]%2!=0{
			l++
		}
		if nums[r]%2==0{
			r--
		}

	}
	return nums
}