package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	oldnums := make([]int, len(nums))
	copy(oldnums, nums)
//	fmt.Println(oldnums)
	newnums := quicksort(nums)
//	fmt.Println(newnums)
//	fmt.Println(oldnums)
	i := 0
	j := len(nums) - 1
//	fmt.Println(j)
	for oldnums[i] == newnums[i] && i < j {
		i++
	}
	for newnums[j] == oldnums[j] && j > i {
		j--
	}
	if i == j {
		return 0
	}
	return j - i + 1

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

func main() {
	arry := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(arry))
}
