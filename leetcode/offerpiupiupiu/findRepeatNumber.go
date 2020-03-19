func findRepeatNumber(nums []int) int {
	m:=make(map[int]int)
	for i,j:=range nums{
		if _,ok:=m[j];ok{
			return j
		}else{
			m[j]=i
		}
	}
    return -1
}