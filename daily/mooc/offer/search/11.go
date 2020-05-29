func minArray(numbers []int) int {
	return deal(numbers,0,len(numbers)-1)
}

func deal(numbers []int,low int,high int)int{
	for numbers[low]>numbers[high]&&low<high{
		m:=(high+low)/2
		if numbers[low]<numbers[m]{
			low=m+1
		}else if numbers[m]<numbers[low]{
			high=m-1
			//可能会有一样的
		}else{
			low++
		}
	}
	return numbers[low]
}
