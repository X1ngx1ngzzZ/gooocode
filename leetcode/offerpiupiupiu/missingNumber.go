func missingNumber(nums []int) int {

	for k,v:=range nums{
		if k!=v{
			return k
		}
}
	return nums[:len(nums)-1]+1
}

func missingNumber(nums []int) int {
	left:=0
	right:=len(nums)
//	mid=left+right
	for left<right{
		mid=(left+right)>>1
		if nums[mid]!=mid{
			right=mid
		}else{
			left=mid+1
		}
	}
	return left
}