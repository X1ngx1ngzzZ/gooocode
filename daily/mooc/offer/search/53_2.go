func missingNumber(nums []int) int {
	low,high:=0,len(nums)
	for high>low{
		m:=(high+low)/2
		if m-=nums[m]{
			low=m
		}else{
			high=m
		}
	}
	return high
}