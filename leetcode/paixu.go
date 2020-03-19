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
	fmt.Println(m)
	bubblesort(m)
	//selectsort(m)
	fmt.Println(len(m))
	//InsertSort(m)
	//Shellsort(m)
	fmt.Println(m)

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
			if i != min {
				arry[i], arry[min] = arry[min], arry[i]
			}
		}
	}
	return arry
}

func InsertSort(arry []int) []int {

	for i := 1; i < len(arry); i++ {
		flag := arry[i]
		key := i - 1
		//不能缺这个Key>=0
		for key >= 0 && flag < arry[key] {
			arry[key+1] = arry[key]
			key--
		}
		if key+1 != i {
			arry[key+1] = flag
		}

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
