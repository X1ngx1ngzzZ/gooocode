package main

import "fmt"
//找出数组中重复的数
func main() {
	s := []int{1, 2, 3, 4, 5, 0, 2}
	//	fmt.Println(dup1(s))
	//	fmt.Println(dup2(s))
	fmt.Println(dup3(s))
	dup3(s)
}

//第一种方法
func dup1(s []int) int {
	if s == nil {
		return -1
	}
	//如果有小于n-1的数，这会返回一个空切片
	s1 := bubblesort(s)
	//要考虑到溢出
	//range的时候k从0开始
	for k, v := range s1 {
		if k == len(s1)-1 {
			return -1
		}
		if v == s[k+1] {
			return v
		}
		//这行仅仅是为了看range从几开始
		fmt.Println(k, v)
	}
	return -1
}

//冒泡排序
func bubblesort(arry []int) []int {
	if arry == nil {
		return nil
	}
	//设立标志位
	flag := true
	//flag为false说明没有交换，直接退出
	for i := 0; i < len(arry) && flag; i++ {
		flag = false
		for j := len(arry) - 2; j >= i; j-- {
			//判断切片里是否有大于N-1的数
			if arry[j+1] > len(arry)-1 {
				return nil
			}
			if arry[j] > arry[j+1] {
				arry[j], arry[j+1] = arry[j+1], arry[j]
				flag = true
			}
		}

	}
	return arry
}

//第二种方法
func dup2(arry []int) int {
	//判断特殊情况，数组为空，传入的数组有大于n-1的数，和没有重复的
	//有个没有判断：传入的数组没有
	if arry == nil {
		return -1
	}
	m := make(map[int]int, len(arry))
	for k, v := range arry {
		//判断map中是否存在
		if _, ok := m[v]; ok {
			return v
			//如果没有就放到map里
		} else {
			m[v] = k
		}
	}
	return -1
}

//第三种方法,这个是真的吊，但是要思路清晰

func dup3(arry []int) int {
	if arry == nil {
		return -1
	}
	for k, v := range arry {
		for k != v {
			if v == arry[v] {
				return v
			} else {
				v, *arry[v] = arry[v], v
			}
		}
	}
	return -1
}
