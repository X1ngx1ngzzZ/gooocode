func moveZeroes(nums []int)  {
	i:=0
	j:=len(nums)-1
	for {
		if i>=j{
			break
		}
		if nums[i]==0{
			nums=append(nums[:i],nums[i+1:]...)
			nums=append(nums,0)
			j--
		}else{
			i++
		}
	}

}