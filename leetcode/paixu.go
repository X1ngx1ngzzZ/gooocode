package main

import "fmt"

/*
type sq struct {
	r      []int
	length int
}
*/

func main() {
	m := []int{109, 137, 49, 190, 87}
	//fmt.Println(m)
	//fmt.Println(bubblesort(m))

	//fmt.Println(len(m))
	//InsertSort(m)
	//Shellsort(m)
	//fmt.Println(m)
	//fmt.Println(mergeSort(m))
	//fmt.Println(selectSortt(m))
	fmt.Println(ShellSort(m))
}

/*
func swap(l *sq, i int, j int) {
	sq.r[i], sq.r[j] = sq.r[j], sq.r[i]
}
*/
func simplebubblesort(arry []int) []int {
	for i := 0; i < len(arry); i++ {
		for j := i + 1; j < len(arry); j++ {
			if arry[i] > arry[j] {
				arry[i], arry[j] = arry[j], arry[i]
			}

		}
	}
	return arry
}

func bubblesort(arry []int) []int {
	flag := true
	for i := 0; i < len(arry) && flag; i++ {
		flag = false
		for j := len(arry) - 2; j >= i; j-- {
			if arry[j] > arry[j+1] {
				arry[j], arry[j+1] = arry[j+1], arry[j]
				flag = true
			}
		}
	}
	return arry
}

/*
func swap(a int, b int) {

}
*/

func selectsort(arry []int) []int {
	for i := 0; i < len(arry); i++ {
		min := i
		for j := i + 1; j < len(arry); j++ {
			if arry[min] > arry[j] {
				min = j
			}
		}
		arry[i], arry[min] = arry[min], arry[i]
	}
	return arry
}
//
func InsertSort(arry []int) []int {
	//109, 137, 49, 190, 87
	for i := 1; i < len(arry); i++ {
		key := i - 1
		for key >= 0 && arry[i] < arry[key] {
			key--
		}
			arry[i],arry[key+1]=arry[key+1],arry[i]
			fmt.Println(arry)
	}
	return arry
}

//错误的sheelsort,只进行了交换没有进行直接插入
func Shellsort(arry []int) []int {
	increment := len(arry)
	for increment > 1 {
		increment = increment/3 + 1
		for i := increment; i < len(arry); i++ {
			if arry[i] < arry[i-increment] {
				arry[i], arry[i-increment] = arry[i-increment], arry[i]
			}
		}

	}

	return arry

}

//正确的shellsort
func ShellSort1(slice []int) {
	var i, j int
	increment := len(slice)
	for {
		//这块是增量序列
		increment = increment/3 + 1
		for i = increment + 1; i < len(slice); i++ {
			//满足下面的if就需要把slice[i]插入有序增量子表
			if slice[i] < slice[i-increment] {
				slice[0] = slice[i]
				for j = i - increment; j > 0 && slice[0] < slice[j]; j = j - increment {
					slice[j+increment] = slice[j]
				}
				slice[j+increment] = slice[0]
			}
		}
		if increment <= 1 {
			break
		}
	}
}

/*
func Heapsort(arry []int) {
	l := len(arry)

}


func buildMaxHeap(arry []int,l int){
	-
}
*/

//快排
func Qicksort(arry []int) []int {
	//由于返回值不同，这里进行递归调用。这里注意切片溢出
	return Qsort(arry, 0, len(arry)-1)
}

func Qsort(arry []int, low int, high int) []int {
	var flag int
	//当左边计数器小于右边的时候，进行func deal
	if low < high {
		//把low和high碰到一起的那个位置返回来
		flag = deal(arry, low, high)

		//对左半部分进行排序
		Qsort(arry, low, flag-1)
		//右半部分排序
		Qsort(arry, flag+1, high)

	}
	return arry
}

//真正排序
/*
func deal(arry []int, low int, high int) int {
	m := high / 2
	choose(arry, low, m, high)
	//设置P为中间的值，让左边比他小，右边比他大，这里设置了low的位置为初始数字，所以从high位开始找
	p := arry[low]
	//low和high碰到的时候就停止
	for low < high {
		for {
			//high位置的大于P就不管，继续往前找，否则就交换low和high的位置
			if arry[high] >= p && low < high {
				high--
			} else {
				arry[low], arry[high] = arry[high], arry[low]
				break
			}
		}
		for {
			//low位小于P就继续向后找，否则交换
			if arry[low] <= p && low < high {
				low++
			} else {
				arry[low], arry[high] = arry[high], arry[low]
				break
			}
		}
	}
	return low
}

func choose(arry []int, low, m, high int) {
	if arry[low] > arry[high] {
		arry[low], arry[high] = arry[high], arry[low]
	} else if arry[m] > arry[high] {
		arry[m], arry[high] = arry[high], arry[m]
	} else if arry[m] > arry[low] {
		arry[m], arry[low] = arry[low], arry[m]
	}
}
*/

//func quicksort(arry []int) []int {
//	//传入低位和高位
//	return sort(arry, 0, len(arry)-1)
//}
//
//func sort(arry []int, low int, high int) []int {
//	var flag int
//	if low < high {
//		flag = deal(arry, low, high)
//		sort(arry, low, flag-1)
//		sort(arry, flag+1, high)
//	}
//	//排完把数组返回去
//	return arry
//}
//
//func deal(arry []int, low int, high int) int {
//	value := arry[low]
//	for low < high {
//		for low < high {
//			if arry[high] >= value {
//				high--
//			} else {
//				arry[low], arry[high] = arry[high], arry[low]
//				break
//			}
//		}
//		for low < high {
//			if arry[low] <= value {
//				low++
//			} else {
//				arry[low], arry[high] = arry[high], arry[low]
//				break
//			}
//		}
//	}
//	return low
//}
//

//2020 7 27 快速排序
func quicksort(arry []int)[]int{
	return quick(arry,0,len(arry)-1)
}

func quick(arry []int,low,high int)[]int{
	var flag int
	if low<high {
		flag = deal(arry, low, high)
		quick(arry, low, flag-1)
		quick(arry, flag+1, high)
	}
	return arry
}


func deal(arry []int,low,high int)int{
	v:=arry[high]
	//如果v取low值，那么就要从high开始
	//取high，就要从low开始
	for low < high {
		if arry[low] <= v {
			low++
		} else {
			arry[low], arry[high] = arry[high], arry[low]
			break
		}
	}

	for low<high {
		for low < high {
			if arry[high] >= v {
				high--
			} else {
				arry[low], arry[high] = arry[high], arry[low]
				break
			}
		}


	}
	return low
}

//2020728 冒泡排序
func bubble(arry []int)[]int{
	for i:=0;i<len(arry);i++{
		for j:=i+1;j<len(arry);j++{
			if arry[i]>arry[j]{
				arry[i],arry[j]=arry[j],arry[i]
			}
		}
	}
	return arry
}

//20200807 归并排序
func mergeSort(arry []int)[]int{
	if len(arry)<2{
		return arry
	}
	mid :=len(arry)/2
	left:=mergeSort(arry[:mid])
	right:=mergeSort(arry[mid:])
	result:=merge(left,right)
	return result
}

func merge(left,right []int)[]int {
	var result []int
	leftIndex := 0
	rightIndex := 0
	//取左边数组遍历，看左边数组的每个数位置
	for leftIndex < len(left) {
		//右边数组先完了
		if rightIndex == len(right) {
			result = append(result, left[leftIndex:]...)
			break
		}
		if left[leftIndex] <= right[rightIndex]{
			result = append(result, left[leftIndex])
		leftIndex++
		}else{
		result = append(result, right[rightIndex])
		rightIndex++
		}
	}
	//左边数组先完了
	result=append(result,right[rightIndex:]...)
	return result
}

//20200807
func selectSortt(arry []int)[]int{
	for i:=0;i<len(arry);i++{
		min:=i
		for j:=i+1;j<len(arry);j++{
			if arry[min]>arry[j]{
				min=j
			}
		}
		if i!=min {
			arry[i], arry[min] = arry[min], arry[i]
		}
	}
	return arry
}

//20200809 shellSort
func ShellSort(a []int)[]int{
	n := len(a)
	h := 1
	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				swap(a, j, j-h)
			}
		}
		h /= 3
	}
	return a
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}


