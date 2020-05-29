package main

//滑动窗口，无法适应负数
func sum(i, j int, arry []int) int {
	res := 0
	for i <= j {
		res += arry[i]
		i++
	}
	return res
}

func subarraySum(nums []int, k int) int {
	i := 0
	j := 1
	res := 0
	if nums[i] == k {
		res++
	}
	for j < len(nums) {
		if sum(i, j, nums) < k {
			j++
		} else if sum(i, j, nums) > k {
			i++
		} else if sum(i, j, nums) == k {
			res++
			j++
		}
	}
	return res
}

//map
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	count := 0
	for _, v := range nums {
		sum += v
		if _, ok := m[sum-k]; ok {
			count += m[sum-k]
		}
		m[sum]++
	}
	return count
}
