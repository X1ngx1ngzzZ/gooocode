package main

import (
	"fmt"
)

func main() {
	m := 4
	n := 6
	//	i := 0
	a := area(m, n)
	//收集无人机的信息成为一个切片
	fmt.Println(a)
	st := []string{"abcde", "defgh", "ghisa", "sdadf", "sdada", "dsadx", "qwert", "tyuio"}
	//st := make([]string, 8)
	//把无人机带回来的字符串输入进去
	//最好放在一个输入函数里
	/*
		for i < 8 {
			fmt.Println("请输入第", i+1, "个字符")
			fmt.Scanln(&st[i])
			i++
		}
	*/
	result := deal(a, st)
	fmt.Println(result)

	//	fmt.Println(transfer("abcde"))

}

//建立二维数组
func area(m, n int) [][]string {
	a := make([][]string, m)
	for k, _ := range a {
		a[k] = make([]string, n)
	}
	return a
}

//选取有用字符串的函数
func transfer(s string) []string {
	slice := []string{}
	for k, v := range s {
		if k%2 == 0 {
			slice = append(slice, string(v))
		}
	}
	return slice

}

//真正的处理函数
func deal(a [][]string, st []string) [][]string {
	i := 0
	j := 1
	k := 0
	var res []string
	//向右
	for {
		if j < 6 {
			res = transfer(st[k])
			a[i][j-1] = res[0]
			a[i][j] = res[1]
			a[i][j+1] = res[2]
			j += 3
			k++
		} else {
			j -= 3
			i++
			break
		}
	}
	for {
		if i < 4 {
			res = transfer(st[k])
			a[i][j-1] = res[0]
			a[i][j] = res[1]
			a[i][j+1] = res[2]
			i++
			k++
		} else {
			i -= 1
			j -= 3
			break
		}
	}
	//向左
	for {
		if j > 0 {
			res = transfer(st[k])
			a[i][j-1] = res[0]
			a[i][j] = res[1]
			a[i][j+1] = res[2]
			j -= 3
			k++
		} else {
			i--
			j += 3
			break
		}
	}
	//向上
	for {
		if i > 0 {
			res = transfer(st[k])
			a[i][j-1] = res[0]
			a[i][j] = res[1]
			a[i][j+1] = res[2]
			i--
			k++
		} else {
			break
		}
	}
	return a
}
