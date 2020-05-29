func search(nums []int, target int) int {
	l := len(nums)
    if l == 0 {
        return 0
    }
	return deal(nums,0,l-1,target)
}

func deal(nums []int,low,high,target int)int{
	var right int
	var left int
	//搜索右边

	for high>=low{
		m:=(low+high)/2
		if nums[m]<=target{
			low=m+1
		}else{
			high=m-1
		}
	}
	right=low
	high=len(nums)-1
	low=0
	//左边界
	
	for high>=low{
		m:=(low+high)/2
		if nums[m]>=target{
			
			high=m-1
			
		}else{
			low=m+1
		}
	}
	left=high
	return  right-left-1
}



func search(nums []int, target int) int {
	if len(nums) == 0 {
        return 0
	}
	var right int
	var left int
	low,high:=0,len(nums)-1
	for high>=low{
		m:=(low+right)/2
		if nums[m]>=target{
			
			high=m-1
			
		}else{
			low=m+1
		}
	}
	left=high
	low,high:=0,len(nums)-1
	for high>=low{
		m:=(low+high)/2
		if nums[m]<=target{
			low=m+1
		}else{
			high=m-1
		}
	}
	right=low
	return  right-left-1
}

