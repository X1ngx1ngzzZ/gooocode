func search(nums []int, target int) int {
	i:=0
	for _,v:=range nums{
		if v==target{
			i++
		}
	}
	return i
}

func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
    
	left, right := 0, size-1

	for left < right {
		mid := left + (right-left)/2  // 寻找左边界
		if nums[mid] < target { 
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return 0
	}

	pleft := left  // 保存 left 的位置
	right = size - 1  // right 重新初始化

	for left < right {
		mid := left + (right-left)/2 + 1 // 寻找右边界
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return right - pleft + 1

