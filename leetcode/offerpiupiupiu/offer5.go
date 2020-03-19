// 替换空格
package main

import (
	"fmt"
)

func main() {
	s := "we are happy"
	fmt.Println(replacespace1(s))
	fmt.Println(replacespace2(s))
}

//不用多余空间，直接在本字符串上进行

func replacespace1(s string) string {
	if s == "" {
		return s
	}
	//遍历字符串看有多少空格
	space := 0
	for _, v := range s {
		//这里很关键，空字符和空格字符是不一样的，这里是空格字符所以一定要在引号间加空格
		if v == ' ' {
			space++
		}
	}
	oldline := len(s) - 1
	newline := oldline + 2*space
	//
	//字符串转切片,然后增加长度
	s1 := []rune(s)
	for i := 0; i < 2*space; i++ {
		s1 = append(s1, ' ')
	}

	for {
		if oldline < 0 {
			break
		}
		if s1[oldline] == ' ' {
			//用到了copy函数
			copy(s1[newline-2:newline+1], []rune{'%', '2', '0'})
			newline -= 3
		} else {
			s1[newline] = s1[oldline]
			newline -= 1
		}
		oldline -= 1
	}
	return string(s1)

}

//用新的空间，拷贝过去
func replacespace2(s string) string {
	news := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			news += "%20"
		} else {
			//这里比较关键，虽然字符串的底层是用切片实现的，但是还是不能直接加。
			news += string(s[i])
		}
	}
	return news
}
