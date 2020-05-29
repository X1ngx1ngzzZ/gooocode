//算法有问题，这不能直接找KV关系，因为原来存在的会向后推
func findDisappearedNumbers(nums []int) []int {
	newnums:=quicksort(nums)
	res:=[]int{}
	for k,v:=range newnums{
		if k+1!=v{
			res=append(res,k+1)
		}
	}
	return res

}


//使用map

func findDisappearedNumbers(nums []int) []int {
	m:=make(map[int]int,len(nums))
	res:=[]int{}
	for _,v:=range nums{
		m[v]++
	}
	for i:=1;i<=len(nums);i++{
		if _,ok:=m[i];!ok{
			res=append(res,i)
		}
	}
	return res
}








//快排
func quicksort(arry []int) []int {
	//传入低位和高位
	return sort(arry, 0, len(arry)-1)
}

func sort(arry []int, low int, high int) []int {
	var flag int
	if low < high {
		flag = deal(arry, low, high)
		sort(arry, low, flag-1)
		sort(arry, flag+1, high)
	}
	//排完把数组返回去
	return arry
}

func deal(arry []int, low int, high int) int {
	value := arry[low]
	for low < high {
		for low < high {
			if arry[high] >= value {
				high--
			} else {
				arry[low], arry[high] = arry[high], arry[low]
				break
			}
		}
		for low < high {
			if arry[low] <= value {
				low++
			} else {
				arry[low], arry[high] = arry[high], arry[low]
				break
			}
		}
	}
	return low
}