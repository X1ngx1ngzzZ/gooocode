func sortColors(nums []int) {
	var i, j int
	k := len(nums) - 1
    i,j =0,0
	for j <= k {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
			i++
		} else if nums[j] == 2 {
			nums[k], nums[j] = nums[j], nums[k]
			k--
		}else if nums[j]==1{
			j++
		}
	}

}
