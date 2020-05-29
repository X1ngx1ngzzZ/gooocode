func maxSubArray(nums []int) int {
	max:=make([]int, len(nums))
	max[0]=nums[0]
	flag:=nums[0]
	for i:=1;i<len(nums);i++{
		if max[i-1]>0{
			max[i]=max[i-1]+nums[i]
			flag=maxValue(flag,max[i])
		}else{
			max[i]=nums[i]
			flag=maxValue(flag,max[i])
		}
	}
	return flag
}

func maxValue(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}