func singleNumbers(nums []int) []int {
    var ret []int
    md := make(map[int]int)
    for _, v := range nums {
        md[v]++
    }
    for k, v := range md {
        if v == 1 {
            ret = append(ret, k)
        }
    }
    return ret
}

func singleNumbers(nums []int)[]int{
	res:=[]int{}
	m:=make(map[int]int)
	for _,v:=range nums{
		m[v]++
	}
	for k,v:=range m{
		if v==1{
			res=append(res,k)
		}
	}
	return res
}