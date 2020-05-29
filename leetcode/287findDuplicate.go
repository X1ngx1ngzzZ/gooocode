//1不能更改原数组（假设数组是只读的）。
//2只能使用额外的 O(1) 的空间。
//3时间复杂度小于 O(n2) 。
//4数组中只有一个重复的数字，但它可能不止重复出现一次

//用 map
//违反2
func findDuplicate(nums []int) int {
	m:=make(map[int]int)
	for _,v:=range nums{
		m[v]++
	}
	for k,v:=range m{
		if v>=2{
			return k
		}
	}
	return -1
}

//弗洛伊德
func findDuplicate(nums []int) int {
	fast:=0
	slow:=0
	//找交点
	for {
		slow=nums[slow]
		fast=nums[fast]
		fast=nums[fast]
		if fast==slow{
			break
		}
	}
	//找入口点
	fast=0
	for{
		fast=nums[fast]
		slow=nums[slow]
		if slow==fast{
			return slow
		}
	}

}