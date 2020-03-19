//二维数组中的查找
package main

import "fmt"

func main() {
	//建立二维数组
	//
	//建立空切片进行测试
	k := make([][]int, 0)
	s := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15}}
	fmt.Println(search(s, 3))
	fmt.Println(search(k, 3))
}

func search(arry [][]int, target int) bool {

	if arry == nil {
		return false
	}

	i := len(arry) - 1
	j := 0
	//一定要注意切片的溢出问题
	for i >= 0 && j <= len(arry)-1 {
		if arry[i][j] < target {
			j++
		} else if arry[i][j] > target {
			i--
		} else if arry[i][j] == target {
			return true
		}
	}
	return false
}
