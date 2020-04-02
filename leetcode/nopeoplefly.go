package main

import (
	"fmt"
)

func main() {
	counts := []string{"GGRF/", "RRGGG", "GRFGR"}
	result := store(counts)
	for _, v := range result {
		fmt.Println(v)
	}
}

//利用大小建立一个二维矩阵（二维数组）
/*
func make(m, n int) [][]int {
		field := make([][]int, m)
		for k, _ := range field {
			field[k] = make([]int, n)
		}
	return field
}
*/
//利用输入的字符串建立一个字符串切片进行存储
/*
func makeslice(str string)[]string{
	res:=[]string
	res = append(res,str)
	return res
}
*/
//把字符串切片的值进行筛选，选出自己需要的
func store(str []string) [][]string {
	var res [][]string
	var temp []string
	for i := 0; i < len(str); i++ {
		temp = transfer(str[i])
		res = append(res, temp)
	}
	return res
}

//返回一个竖着存储信息的字符串
func transfer(str string) []string {
	res := []string{string(str[0]), string(str[1]), string(str[4])}
	return res
}
